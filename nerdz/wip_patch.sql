BEGIN;

drop table if exists oauth2_clients cascade;
create table oauth2_clients(
    id bigserial not null primary key,
    name varchar(100) not null unique,
    secret text not null unique,
    redirect_uri varchar(350) not null,
    user_id bigint not null references users(counter) on delete cascade
);

drop table if exists oauth2_authorize cascade;
create table oauth2_authorize (
    id bigserial not null primary key, -- surrogated id
    code text not null unique, --logical id
    client_id bigint not null references oauth2_clients(id) on delete cascade,
    created_At timestamp without time zone not null default (now() at time zone 'utc'),
    expires_in bigint not null,
    state text not null,
    scope text not null,
    redirect_uri varchar(350) not null,
    user_id bigint not null references users(counter) on delete cascade
);

drop table if exists oauth2_refresh cascade;
create table oauth2_refresh(
    id bigserial not null primary key, -- surrogated id
    token text not null unique -- logical id
);

drop table if exists oauth2_access cascade;
create table oauth2_access(
    id bigserial not null primary key, -- surrogated id
    client_id bigint not null references oauth2_clients(id) on delete cascade,
    access_token text not null unique, -- logical id
    created_At timestamp without time zone not null default (now() at time zone 'utc'),
    expires_in bigint not null,
    redirect_uri varchar(350) not null,
    oauth2_authorize_id bigint NULL references oauth2_authorize(id) on delete cascade,
    oauth2_access_id bigint NULL references oauth2_access(id) on delete cascade,
    refresh_token_id bigint NULL references oauth2_refresh(id) on delete cascade,
    scope text not null,
    user_id bigint not null references users(counter) on delete cascade
);

drop view if exists messages cascade;

CREATE OR REPLACE FUNCTION after_delete_blacklist() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

    BEGIN
    
        DELETE FROM "posts_no_notify" WHERE "user" = OLD."to" AND (
            "hpid" IN (
            
                SELECT "hpid"  FROM "posts" WHERE "from" = OLD."to" AND "to" = OLD."from"
                
            ) OR "hpid" IN (
            
                SELECT "hpid"  FROM "comments" WHERE "from" = OLD."to" AND "to" = OLD."from"
                
            )
        );
        
        RETURN OLD;
        
    END

$$;

CREATE OR REPLACE FUNCTION after_delete_user() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
    insert into deleted_users(counter, username) values(OLD.counter, OLD.username);
    RETURN NULL;
    -- if the user gives a motivation, the upper level might update this row
end $$;



--
-- Name: after_insert_blacklist(); Type: FUNCTION; Schema: public; 
--

CREATE OR REPLACE FUNCTION after_insert_blacklist() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE r RECORD;
BEGIN
    INSERT INTO posts_no_notify("user","hpid")
    (
        SELECT NEW."from", "hpid" FROM "posts" WHERE "to" = NEW."to" OR "from" = NEW."to" -- posts made by the blacklisted user and post on his board
            UNION DISTINCT
        SELECT NEW."from", "hpid" FROM "comments" WHERE "from" = NEW."to" OR "to" = NEW."to" -- comments made by blacklisted user on others and his board
    )
    EXCEPT -- except existing ones
    (
        SELECT NEW."from", "hpid" FROM "posts_no_notify" WHERE "user" = NEW."from"
    );

    INSERT INTO groups_posts_no_notify("user","hpid")
    (
        (
            SELECT NEW."from", "hpid" FROM "groups_posts" WHERE "from" = NEW."to" -- posts made by the blacklisted user in every project
                UNION DISTINCT
            SELECT NEW."from", "hpid" FROM "groups_comments" WHERE "from" = NEW."to" -- comments made by the blacklisted user in every project
        )
        EXCEPT -- except existing ones
        (
            SELECT NEW."from", "hpid" FROM "groups_posts_no_notify" WHERE "user" = NEW."from"
        )
    );
    

    FOR r IN (SELECT "to" FROM "groups_owners" WHERE "from" = NEW."from")
    LOOP
        -- remove from my groups members
        DELETE FROM "groups_members" WHERE "from" = NEW."to" AND "to" = r."to";
    END LOOP;
    
    -- remove from followers
    DELETE FROM "followers" WHERE ("from" = NEW."from" AND "to" = NEW."to");

    -- remove pms
    DELETE FROM "pms" WHERE ("from" = NEW."from" AND "to" = NEW."to") OR ("to" = NEW."from" AND "from" = NEW."to");

    -- remove from mentions
    DELETE FROM "mentions" WHERE ("from"= NEW."from" AND "to" = NEW."to") OR ("to" = NEW."from" AND "from" = NEW."to");

    RETURN NULL;
END $$;

CREATE OR REPLACE FUNCTION after_insert_group_post() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
 BEGIN
     WITH to_notify("user") AS (
         (
             -- members
             SELECT "from" FROM "groups_members" WHERE "to" = NEW."to"
                 UNION DISTINCT
             --followers
             SELECT "from" FROM "groups_followers" WHERE "to" = NEW."to"
                 UNION DISTINCT
             SELECT "from"  FROM "groups_owners" WHERE "to" = NEW."to"
         )
         EXCEPT
         (
             -- blacklist
             SELECT "from" AS "user" FROM "blacklist" WHERE "to" = NEW."from"
                 UNION DISTINCT
             SELECT "to" AS "user" FROM "blacklist" WHERE "from" = NEW."from"
                 UNION DISTINCT
             SELECT NEW."from" -- I shouldn't be notified about my new post
         )
     )

     INSERT INTO "groups_notify"("from", "to", "time", "hpid") (
         SELECT NEW."to", "user", NEW."time", NEW."hpid" FROM to_notify
     );

     PERFORM hashtag(NEW.message, NEW.hpid, true, NEW.from, NEW.time);
     PERFORM mention(NEW."from", NEW.message, NEW.hpid, true);
     RETURN NULL;
 END $$;

CREATE OR REPLACE FUNCTION after_insert_user() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        INSERT INTO "profiles"(counter) VALUES(NEW.counter);
        RETURN NULL;
    END $$;

CREATE OR REPLACE FUNCTION after_insert_user_post() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    begin
        IF NEW."from" <> NEW."to" THEN
         insert into posts_notify("from", "to", "hpid", "time") values(NEW."from", NEW."to", NEW."hpid", NEW."time");
        END IF;
        PERFORM hashtag(NEW.message, NEW.hpid, false, NEW.from, NEW.time);
        PERFORM mention(NEW."from", NEW.message, NEW.hpid, false);
        return null;
    end $$;



--
-- Name: after_update_userame(); Type: FUNCTION; Schema: public; 
--

CREATE OR REPLACE FUNCTION after_update_userame() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- create news
    insert into posts("from","to","message")
    SELECT counter, counter,
    OLD.username || ' %%12now is34%% [user]' || NEW.username || '[/user]' FROM special_users WHERE "role" = 'GLOBAL_NEWS';

    RETURN NULL;
END $$;

CREATE OR REPLACE FUNCTION before_delete_user() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        UPDATE "comments" SET "from" = (SELECT "counter" FROM "special_users" WHERE "role" = 'DELETED') WHERE "from" = OLD.counter;
        UPDATE "posts" SET "from" = (SELECT "counter" FROM "special_users" WHERE "role" = 'DELETED') WHERE "from" = OLD.counter;

        UPDATE "groups_comments" SET "from" = (SELECT "counter" FROM "special_users" WHERE "role" = 'DELETED') WHERE "from" = OLD.counter;            
        UPDATE "groups_posts" SET "from" = (SELECT "counter" FROM "special_users" WHERE "role" = 'DELETED') WHERE "from" = OLD.counter;

        PERFORM handle_groups_on_user_delete(OLD.counter);

        RETURN OLD;
    END
$$;

CREATE OR REPLACE FUNCTION before_insert_comment() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE closedPost boolean;
BEGIN
    PERFORM flood_control('"comments"', NEW."from", NEW.message);
    SELECT closed FROM posts INTO closedPost WHERE hpid = NEW.hpid;
    IF closedPost THEN
        RAISE EXCEPTION 'CLOSED_POST';
    END IF;

    SELECT p."to" INTO NEW."to" FROM "posts" p WHERE p.hpid = NEW.hpid;
    PERFORM blacklist_control(NEW."from", NEW."to");

    NEW.message = message_control(NEW.message);

    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION before_insert_comment_thumb() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE postFrom int8;
        tmp record;
BEGIN
    PERFORM flood_control('"comment_thumbs"', NEW."from");

    SELECT T."to", T."from", T."hpid" INTO tmp FROM (SELECT "from", "to", "hpid" FROM "comments" WHERE "hcid" = NEW.hcid) AS T;
    SELECT tmp."from" INTO NEW."to";

    PERFORM blacklist_control(NEW."from", NEW."to"); --blacklisted commenter

    SELECT T."from", T."to" INTO tmp FROM (SELECT p."from", p."to" FROM "posts" p WHERE p.hpid = tmp.hpid) AS T;

    PERFORM blacklist_control(NEW."from", tmp."from"); --blacklisted post creator
    IF tmp."from" <> tmp."to" THEN
        PERFORM blacklist_control(NEW."from", tmp."to"); --blacklisted post destination user
    END IF;

    IF NEW."vote" = 0 THEN
        DELETE FROM "comment_thumbs" WHERE hcid = NEW.hcid AND "from" = NEW."from";
        RETURN NULL;
    END IF;
    
    WITH new_values (hcid, "from", vote) AS (
            VALUES(NEW."hcid", NEW."from", NEW."vote")
        ),
        upsert AS (
            UPDATE "comment_thumbs" AS m
            SET vote = nv.vote
            FROM new_values AS nv
            WHERE m.hcid = nv.hcid AND m."from" = nv."from"
            RETURNING m.*
       )

       SELECT "vote" INTO NEW."vote"
       FROM new_values
       WHERE NOT EXISTS (
           SELECT 1
           FROM upsert AS up
           WHERE up.hcid = new_values.hcid AND up."from" = new_values."from"
      );

    IF NEW."vote" IS NULL THEN -- updated previous vote
        RETURN NULL; --no need to insert new value
    END IF;
    
    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION before_insert_follower() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    PERFORM flood_control('"followers"', NEW."from");
    IF NEW."from" = NEW."to" THEN
        RAISE EXCEPTION 'CANT_FOLLOW_YOURSELF';
    END IF;
    PERFORM blacklist_control(NEW."from", NEW."to");
    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION before_insert_group_post_lurker() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE tmp RECORD;
BEGIN
    PERFORM flood_control('"groups_lurkers"', NEW."from");

    SELECT T."to", T."from" INTO tmp FROM (SELECT "to", "from" FROM "groups_posts" WHERE "hpid" = NEW.hpid) AS T;

    SELECT tmp."to" INTO NEW."to";

    PERFORM blacklist_control(NEW."from", tmp."from"); --blacklisted post creator

    IF NEW."from" IN ( SELECT "from" FROM "groups_comments" WHERE hpid = NEW.hpid ) THEN
        RAISE EXCEPTION 'CANT_LURK_IF_POSTED';
    END IF;
    
    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION before_insert_groups_comment() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE postFrom int8;
        closedPost boolean;
BEGIN
    PERFORM flood_control('"groups_comments"', NEW."from", NEW.message);

    SELECT closed FROM groups_posts INTO closedPost WHERE hpid = NEW.hpid;
    IF closedPost THEN
        RAISE EXCEPTION 'CLOSED_POST';
    END IF;

    SELECT p."to" INTO NEW."to" FROM "groups_posts" p WHERE p.hpid = NEW.hpid;

    NEW.message = message_control(NEW.message);


    SELECT T."from" INTO postFrom FROM (SELECT "from" FROM "groups_posts" WHERE hpid = NEW.hpid) AS T;
    PERFORM blacklist_control(NEW."from", postFrom); --blacklisted post creator

    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION before_insert_groups_comment_thumb() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE tmp record;
        postFrom int8;
BEGIN
    PERFORM flood_control('"groups_comment_thumbs"', NEW."from");

    SELECT T."hpid", T."from", T."to" INTO tmp FROM (SELECT "hpid", "from","to" FROM "groups_comments" WHERE "hcid" = NEW.hcid) AS T;
    SELECT tmp."from" INTO NEW."to";

    PERFORM blacklist_control(NEW."from", NEW."to"); --blacklisted commenter

    SELECT T."from" INTO postFrom FROM (SELECT p."from" FROM "groups_posts" p WHERE p.hpid = tmp.hpid) AS T;

    PERFORM blacklist_control(NEW."from", postFrom); --blacklisted post creator

    IF NEW."vote" = 0 THEN
        DELETE FROM "groups_comment_thumbs" WHERE hcid = NEW.hcid AND "from" = NEW."from";
        RETURN NULL;
    END IF;

    WITH new_values (hcid, "from", vote) AS (
            VALUES(NEW."hcid", NEW."from", NEW."vote")
        ),
        upsert AS (
            UPDATE "groups_comment_thumbs" AS m
            SET vote = nv.vote
            FROM new_values AS nv
            WHERE m.hcid = nv.hcid AND m."from" = nv."from"
            RETURNING m.*
       )

       SELECT "vote" INTO NEW."vote"
       FROM new_values
       WHERE NOT EXISTS (
           SELECT 1
           FROM upsert AS up
           WHERE up.hcid = new_values.hcid AND up."from" = new_values."from"
      );

    IF NEW."vote" IS NULL THEN -- updated previous vote
        RETURN NULL; --no need to insert new value
    END IF;
    
    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION before_insert_groups_follower() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE group_owner int8;
BEGIN
    PERFORM flood_control('"groups_followers"', NEW."from");
    SELECT "from" INTO group_owner FROM "groups_owners" WHERE "to" = NEW."to";
    PERFORM blacklist_control(group_owner, NEW."from");
    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION before_insert_groups_member() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE group_owner int8;
BEGIN
    SELECT "from" INTO group_owner FROM "groups_owners" WHERE "to" = NEW."to";
    PERFORM blacklist_control(group_owner, NEW."from");
    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION before_insert_groups_thumb() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE  tmp record;
BEGIN
    PERFORM flood_control('"groups_thumbs"', NEW."from");

    SELECT T."to", T."from" INTO tmp
    FROM (SELECT "to", "from" FROM "groups_posts" WHERE "hpid" = NEW.hpid) AS T;

    SELECT tmp."from" INTO NEW."to";

    PERFORM blacklist_control(NEW."from", NEW."to"); -- blacklisted post creator

    IF NEW."vote" = 0 THEN
        DELETE FROM "groups_thumbs" WHERE hpid = NEW.hpid AND "from" = NEW."from";
        RETURN NULL;
    END IF;

    WITH new_values (hpid, "from", vote) AS (
            VALUES(NEW."hpid", NEW."from", NEW."vote")
        ),
        upsert AS (
            UPDATE "groups_thumbs" AS m
            SET vote = nv.vote
            FROM new_values AS nv
            WHERE m.hpid = nv.hpid AND m."from" = nv."from"
            RETURNING m.*
       )

       SELECT "vote" INTO NEW."vote"
       FROM new_values
       WHERE NOT EXISTS (
           SELECT 1
           FROM upsert AS up
           WHERE up.hpid = new_values.hpid AND up."from" = new_values."from"
      );

    IF NEW."vote" IS NULL THEN -- updated previous vote
        RETURN NULL; --no need to insert new value
    END IF;
    
    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION before_insert_pm() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE myLastMessage RECORD;
BEGIN
    NEW.message = message_control(NEW.message);
    PERFORM flood_control('"pms"', NEW."from", NEW.message);

    IF NEW."from" = NEW."to" THEN
        RAISE EXCEPTION 'CANT_PM_YOURSELF';
    END IF;

    PERFORM blacklist_control(NEW."from", NEW."to");
    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION before_insert_thumb() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE tmp RECORD;
BEGIN
    PERFORM flood_control('"thumbs"', NEW."from");

    SELECT T."to", T."from" INTO tmp FROM (SELECT "to", "from" FROM "posts" WHERE "hpid" = NEW.hpid) AS T;

    SELECT tmp."from" INTO NEW."to";

    PERFORM blacklist_control(NEW."from", NEW."to"); -- can't thumb on blacklisted board
    IF tmp."from" <> tmp."to" THEN
        PERFORM blacklist_control(NEW."from", tmp."from"); -- can't thumbs if post was made by blacklisted user
    END IF;

    IF NEW."vote" = 0 THEN
        DELETE FROM "thumbs" WHERE hpid = NEW.hpid AND "from" = NEW."from";
        RETURN NULL;
    END IF;
   
    WITH new_values (hpid, "from", vote) AS (
            VALUES(NEW."hpid", NEW."from", NEW."vote")
        ),
        upsert AS (
            UPDATE "thumbs" AS m
            SET vote = nv.vote
            FROM new_values AS nv
            WHERE m.hpid = nv.hpid AND m."from" = nv."from"
            RETURNING m.*
       )

       SELECT "vote" INTO NEW."vote"
       FROM new_values
       WHERE NOT EXISTS (
           SELECT 1
           FROM upsert AS up
           WHERE up.hpid = new_values.hpid AND up."from" = new_values."from"
      );

    IF NEW."vote" IS NULL THEN -- updated previous vote
        RETURN NULL; --no need to insert new value
    END IF;
    
    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION before_insert_user_post_lurker() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE tmp RECORD;
BEGIN
    PERFORM flood_control('"lurkers"', NEW."from");

    SELECT T."to", T."from" INTO tmp FROM (SELECT "to", "from" FROM "posts" WHERE "hpid" = NEW.hpid) AS T;

    SELECT tmp."to" INTO NEW."to";

    PERFORM blacklist_control(NEW."from", NEW."to"); -- can't lurk on blacklisted board
    IF tmp."from" <> tmp."to" THEN
        PERFORM blacklist_control(NEW."from", tmp."from"); -- can't lurk if post was made by blacklisted user
    END IF;

    IF NEW."from" IN ( SELECT "from" FROM "comments" WHERE hpid = NEW.hpid ) THEN
        RAISE EXCEPTION 'CANT_LURK_IF_POSTED';
    END IF;
    
    RETURN NEW;
    
END $$;

CREATE OR REPLACE FUNCTION blacklist_control(me bigint, other bigint) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- templates and other implementations must handle exceptions with localized functions
    IF me IN (SELECT "from" FROM blacklist WHERE "to" = other) THEN
        RAISE EXCEPTION 'YOU_BLACKLISTED_THIS_USER';
    END IF;

    IF me IN (SELECT "to" FROM blacklist WHERE "from" = other) THEN
        RAISE EXCEPTION 'YOU_HAVE_BEEN_BLACKLISTED';
    END IF;
END $$;

CREATE OR REPLACE FUNCTION flood_control(tbl regclass, flooder bigint, message text DEFAULT NULL::text) RETURNS void
    LANGUAGE plpgsql
    AS $$
DECLARE now timestamp(0) without time zone;
        lastAction timestamp(0) without time zone;
        interv interval minute to second;
        myLastMessage text;
        postId text;
BEGIN
    EXECUTE 'SELECT MAX("time") FROM ' || tbl || ' WHERE "from" = ' || flooder || ';' INTO lastAction;
    now := (now() at time zone 'utc');

    SELECT time FROM flood_limits WHERE table_name = tbl INTO interv;

    IF now - lastAction < interv THEN
        RAISE EXCEPTION 'FLOOD ~%~', interv - (now - lastAction);
    END IF;

    -- duplicate messagee
    IF message IS NOT NULL AND tbl IN ('comments', 'groups_comments', 'posts', 'groups_posts') THEN
        
        SELECT CASE
           WHEN tbl IN ('comments', 'groups_comments') THEN 'hcid'
           WHEN tbl IN ('posts', 'groups_posts') THEN 'hpid'
           ELSE 'pmid'
        END AS columnName INTO postId;

        EXECUTE 'SELECT "message" FROM ' || tbl || ' WHERE "from" = ' || flooder || ' AND ' || postId || ' = (
            SELECT MAX(' || postId ||') FROM ' || tbl || ' WHERE "from" = ' || flooder || ')' INTO myLastMessage;

        IF myLastMessage = message THEN
            RAISE EXCEPTION 'FLOOD';
        END IF;
    END IF;
END $$;

CREATE OR REPLACE FUNCTION group_comment() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
 BEGIN
     PERFORM hashtag(NEW.message, NEW.hpid, true, NEW.from, NEW.time);
     PERFORM mention(NEW."from", NEW.message, NEW.hpid, true);
     -- edit support
     IF TG_OP = 'UPDATE' THEN
         INSERT INTO groups_comments_revisions(hcid, time, message, rev_no)
         VALUES(OLD.hcid, OLD.time, OLD.message, (
             SELECT COUNT(hcid) + 1 FROM groups_comments_revisions WHERE hcid = OLD.hcid
         ));

          --notify only if it's the last comment in the post
         IF OLD.hcid <> (SELECT MAX(hcid) FROM groups_comments WHERE hpid = NEW.hpid) THEN
             RETURN NULL;
         END IF;
     END IF;


     -- if I commented the post, I stop lurking
     DELETE FROM "groups_lurkers" WHERE "hpid" = NEW."hpid" AND "from" = NEW."from";

     WITH no_notify("user") AS (
         -- blacklist
         (
             SELECT "from" FROM "blacklist" WHERE "to" = NEW."from"
                 UNION
             SELECT "to" FROM "blacklist" WHERE "from" = NEW."from"
         )
         UNION -- users that locked the notifications for all the thread
             SELECT "user" FROM "groups_posts_no_notify" WHERE "hpid" = NEW."hpid"
         UNION -- users that locked notifications from me in this thread
             SELECT "to" FROM "groups_comments_no_notify" WHERE "from" = NEW."from" AND "hpid" = NEW."hpid"
         UNION -- users mentioned in this post (already notified, with the mention)
             SELECT "to" FROM "mentions" WHERE "g_hpid" = NEW.hpid AND to_notify IS TRUE
         UNION
             SELECT NEW."from"
     ),
     to_notify("user") AS (
             SELECT DISTINCT "from" FROM "groups_comments" WHERE "hpid" = NEW."hpid"
         UNION
             SELECT "from" FROM "groups_lurkers" WHERE "hpid" = NEW."hpid"
         UNION
             SELECT "from" FROM "groups_posts" WHERE "hpid" = NEW."hpid"
     ),
     real_notify("user") AS (
         -- avoid to add rows with the same primary key
         SELECT "user" FROM (
             SELECT "user" FROM to_notify
                 EXCEPT
             (
                 SELECT "user" FROM no_notify
              UNION
                 SELECT "to" FROM "groups_comments_notify" WHERE "hpid" = NEW."hpid"
             )
         ) AS T1
     )

     INSERT INTO "groups_comments_notify"("from","to","hpid","time") (
         SELECT NEW."from", "user", NEW."hpid", NEW."time" FROM real_notify
     );

     RETURN NULL;
 END $$;

CREATE OR REPLACE FUNCTION group_comment_edit_control() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE postFrom int8;
BEGIN
    IF OLD.editable IS FALSE THEN
        RAISE EXCEPTION 'NOT_EDITABLE';
    END IF;

    -- update time
    SELECT (now() at time zone 'utc') INTO NEW.time;

    NEW.message = message_control(NEW.message);
    PERFORM flood_control('"groups_comments"', NEW."from", NEW.message);

    SELECT T."from" INTO postFrom FROM (SELECT "from" FROM "groups_posts" WHERE hpid = NEW.hpid) AS T;
    PERFORM blacklist_control(NEW."from", postFrom); --blacklisted post creator

    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION group_interactions(me bigint, grp bigint) RETURNS SETOF record
    LANGUAGE plpgsql
    AS $$
DECLARE tbl text;
        ret record;
        query text;
BEGIN
    FOR tbl IN (SELECT unnest(array['groups_members', 'groups_followers', 'groups_comments', 'groups_comment_thumbs', 'groups_lurkers', 'groups_owners', 'groups_thumbs', 'groups_posts'])) LOOP
        query := interactions_query_builder(tbl, me, grp, true);
        FOR ret IN EXECUTE query LOOP
            RETURN NEXT ret;
        END LOOP;
    END LOOP;
   RETURN;
END $$;

CREATE OR REPLACE FUNCTION group_post_control() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE group_owner int8;
        open_group boolean;
        members int8[];
BEGIN
    NEW.message = message_control(NEW.message);

    IF TG_OP = 'INSERT' THEN -- no flood control on update
        PERFORM flood_control('"groups_posts"', NEW."from", NEW.message);
    END IF;

    SELECT "from" INTO group_owner FROM "groups_owners" WHERE "to" = NEW."to";
    SELECT "open" INTO open_group FROM groups WHERE "counter" = NEW."to";

    IF group_owner <> NEW."from" AND
        (
            open_group IS FALSE AND NEW."from" NOT IN (
                SELECT "from" FROM "groups_members" WHERE "to" = NEW."to" )
        )
    THEN
        RAISE EXCEPTION 'CLOSED_PROJECT';
    END IF;

    IF open_group IS FALSE THEN -- if the group is closed, blacklist works
        PERFORM blacklist_control(NEW."from", group_owner);
    END IF;

    IF TG_OP = 'UPDATE' THEN
        SELECT (now() at time zone 'utc') INTO NEW.time;
    ELSE
        SELECT "pid" INTO NEW.pid FROM (
            SELECT COALESCE( (SELECT "pid" + 1 as "pid" FROM "groups_posts"
            WHERE "to" = NEW."to"
            ORDER BY "hpid" DESC
            FETCH FIRST ROW ONLY), 1) AS "pid"
        ) AS T1;
    END IF;

    IF NEW."from" <> group_owner AND NEW."from" NOT IN (
        SELECT "from" FROM "groups_members" WHERE "to" = NEW."to"
    ) THEN
        SELECT false INTO NEW.news; -- Only owner and members can send news
    END IF;

    -- if to = GLOBAL_NEWS set the news filed to true
    IF NEW."to" = (SELECT counter FROM special_groups where "role" = 'GLOBAL_NEWS') THEN
        SELECT true INTO NEW.news;
    END IF;

    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION groups_post_update() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
 BEGIN
     INSERT INTO groups_posts_revisions(hpid, time, message, rev_no) VALUES(OLD.hpid, OLD.time, OLD.message,
         (SELECT COUNT(hpid) +1 FROM groups_posts_revisions WHERE hpid = OLD.hpid));

     PERFORM hashtag(NEW.message, NEW.hpid, true, NEW.from, NEW.time);
     PERFORM mention(NEW."from", NEW.message, NEW.hpid, true);
     RETURN NULL;
 END $$;

CREATE OR REPLACE FUNCTION handle_groups_on_user_delete(usercounter bigint) RETURNS void
    LANGUAGE plpgsql
    AS $$
declare r RECORD;
newOwner int8;
begin
    FOR r IN SELECT "to" FROM "groups_owners" WHERE "from" = userCounter LOOP
        IF EXISTS (select "from" FROM groups_members where "to" = r."to") THEN
            SELECT gm."from" INTO newowner FROM groups_members gm
            WHERE "to" = r."to" AND "time" = (
                SELECT min(time) FROM groups_members WHERE "to" = r."to"
            );
            
            UPDATE "groups_owners" SET "from" = newOwner, to_notify = TRUE WHERE "to" = r."to";
            DELETE FROM groups_members WHERE "from" = newOwner;
        END IF;
        -- else, the foreing key remains and the group will be dropped
    END LOOP;
END $$;

CREATE OR REPLACE FUNCTION hashtag(message text, hpid bigint, grp boolean, from_u bigint, m_time timestamp without time zone) RETURNS void
    LANGUAGE plpgsql
    AS $$
     declare field text;
             regex text;
BEGIN
     IF grp THEN
         field := 'g_hpid';
     ELSE
         field := 'u_hpid';
     END IF;

     regex = '((?![\d]+[[^\w]+|])[\w]{1,44})';

     message = quote_literal(message);

     EXECUTE '
     insert into posts_classification(' || field || ' , "from", time, tag)
     select distinct ' || hpid ||', ' || from_u || ', ''' || m_time || '''::timestamptz, tmp.matchedTag[1] from (
         -- 1: existing hashtags
        select concat(''{#'', a.matchedTag[1], ''}'')::text[] as matchedTag from (
            select regexp_matches(' || strip_tags(message) || ', ''(?:\s|^|\W)#' || regex || ''', ''gi'')
            as matchedTag
        ) as a
             union distinct -- 2: spoiler
         select concat(''{#'', b.matchedTag[1], ''}'')::text[] from (
             select regexp_matches(' || message || ', ''\[spoiler=' || regex || '\]'', ''gi'')
             as matchedTag
         ) as b
             union distinct -- 3: languages
          select concat(''{#'', c.matchedTag[1], ''}'')::text[] from (
              select regexp_matches(' || message || ', ''\[code=' || regex || '\]'', ''gi'')
             as matchedTag
         ) as c
            union distinct -- 4: languages, short tag
         select concat(''{#'', d.matchedTag[1], ''}'')::text[] from (
              select regexp_matches(' || message || ', ''\[c=' || regex || '\]'', ''gi'')
             as matchedTag
         ) as d
     ) tmp
     where not exists (
        select 1
        from posts_classification p
        where ' || field ||'  = ' || hpid || ' and
            p.tag = tmp.matchedTag[1] and
            p.from = ' || from_u || ' -- store user association with tag even if tag already exists
     )';
END $$;

CREATE OR REPLACE FUNCTION interactions_query_builder(tbl text, me bigint, other bigint, grp boolean) RETURNS text
    LANGUAGE plpgsql
    AS $$
declare ret text;
begin
    ret := 'SELECT ''' || tbl || '''::text';
    IF NOT grp THEN
        ret = ret || ' ,t."from", t."to"';
    END IF;
    ret = ret || ', t."time" ';
    --joins
        IF tbl ILIKE '%comments' OR tbl = 'thumbs' OR tbl = 'groups_thumbs' OR tbl ILIKE '%lurkers'
        THEN

            ret = ret || ' , p."pid", p."to" FROM "' || tbl || '" t INNER JOIN "';
            IF grp THEN
                ret = ret || 'groups_';
            END IF;
            ret = ret || 'posts" p ON p.hpid = t.hpid';

        ELSIF tbl ILIKE '%posts' THEN

            ret = ret || ', "pid", "to" FROM "' || tbl || '" t';

        ELSIF tbl ILIKE '%comment_thumbs' THEN

            ret = ret || ', p."pid", p."to" FROM "';

            IF grp THEN
                ret = ret || 'groups_';
            END IF;

            ret = ret || 'comments" c INNER JOIN "' || tbl || '" t
                ON t.hcid = c.hcid
            INNER JOIN "';

            IF grp THEN
                ret = ret || 'groups_';
            END IF;

            ret = ret || 'posts" p ON p.hpid = c.hpid';

        ELSE
            ret = ret || ', null::int8, null::int8  FROM ' || tbl || ' t ';

        END IF;
    --conditions
    ret = ret || ' WHERE (t."from" = '|| me ||' AND t."to" = '|| other ||')';

    IF NOT grp THEN
        ret = ret || ' OR (t."from" = '|| other ||' AND t."to" = '|| me ||')';
    END IF;

    RETURN ret;
end $$;

--
-- Name: mention(bigint, text, bigint, boolean); Type: FUNCTION; Schema: public; 
--

CREATE OR REPLACE FUNCTION mention(me bigint, message text, hpid bigint, grp boolean) RETURNS void
    LANGUAGE plpgsql
    AS $$
DECLARE field text;
    posts_notify_tbl text;
    comments_notify_tbl text;
    posts_no_notify_tbl text;
    comments_no_notify_tbl text;
    project record;
    owner int8;
    other int8;
    matches text[];
    username text;
    found boolean;
BEGIN
    -- prepare tables
    IF grp THEN
        EXECUTE 'SELECT closed FROM groups_posts WHERE hpid = ' || hpid INTO found;
        IF found THEN
            RETURN;
        END IF;
        posts_notify_tbl = 'groups_notify';
        posts_no_notify_tbl = 'groups_posts_no_notify';

        comments_notify_tbl = 'groups_comments_notify';
        comments_no_notify_tbl = 'groups_comments_no_notify';
    ELSE
        EXECUTE 'SELECT closed FROM posts WHERE hpid = ' || hpid INTO found;
        IF found THEN
            RETURN;
        END IF;
        posts_notify_tbl = 'posts_notify';
        posts_no_notify_tbl = 'posts_no_notify';

        comments_notify_tbl = 'comments_notify';
        comments_no_notify_tbl = 'comments_no_notify';           
    END IF;

    -- extract [user]username[/user]
    message = quote_literal(message);
    FOR matches IN
        EXECUTE 'select regexp_matches(' || message || ',
            ''(?!\[(?:url|code|video|yt|youtube|music|img|twitter)[^\]]*\])\[user\](.+?)\[/user\](?![^\[]*\[\/(?:url|code|video|yt|youtube|music|img|twitter)\])'', ''gi''
        )' LOOP

        username = matches[1];
        -- if username exists
        EXECUTE 'SELECT counter FROM users WHERE LOWER(username) = LOWER(' || quote_literal(username) || ');' INTO other;
        IF other IS NULL OR other = me THEN
            CONTINUE;
        END IF;

        -- check if 'other' is in notfy list.
        -- if it is, continue, since he will receive notification about this post anyway
        EXECUTE 'SELECT ' || other || ' IN (
            (SELECT "to" FROM "' || posts_notify_tbl || '" WHERE hpid = ' || hpid || ')
                UNION
           (SELECT "to" FROM "' || comments_notify_tbl || '" WHERE hpid = ' || hpid || ')
        )' INTO found;

        IF found THEN
            CONTINUE;
        END IF;

        -- check if 'ohter' disabled notification from post hpid, if yes -> skip
        EXECUTE 'SELECT ' || other || ' IN (SELECT "user" FROM "' || posts_no_notify_tbl || '" WHERE hpid = ' || hpid || ')' INTO found;
        IF found THEN
            CONTINUE;
        END IF;

        --check if 'other' disabled notification from 'me' in post hpid, if yes -> skip
        EXECUTE 'SELECT ' || other || ' IN (SELECT "to" FROM "' || comments_no_notify_tbl || '" WHERE hpid = ' || hpid || ' AND "from" = ' || me || ')' INTO found;

        IF found THEN
            CONTINUE;
        END IF;

        -- blacklist control
        BEGIN
            PERFORM blacklist_control(me, other);

            IF grp THEN
                EXECUTE 'SELECT counter, visible
                FROM groups WHERE "counter" = (
                    SELECT "to" FROM groups_posts p WHERE p.hpid = ' || hpid || ');'
                INTO project;

                select "from" INTO owner FROM groups_owners WHERE "to" = project.counter;
                -- other can't access groups if the owner blacklisted him
                PERFORM blacklist_control(owner, other);

                -- if the project is NOT visible and other is not the owner or a member
                IF project.visible IS FALSE AND other NOT IN (
                    SELECT "from" FROM groups_members WHERE "to" = project.counter
                        UNION
                      SELECT owner
                    ) THEN
                    RETURN;
                END IF;
            END IF;

        EXCEPTION
            WHEN OTHERS THEN
                CONTINUE;
        END;

        IF grp THEN
            field := 'g_hpid';
        ELSE
            field := 'u_hpid';
        END IF;

        -- if here and mentions does not exists, insert
        EXECUTE 'INSERT INTO mentions(' || field || ' , "from", "to")
        SELECT ' || hpid || ', ' || me || ', '|| other ||'
        WHERE NOT EXISTS (
            SELECT 1 FROM mentions
            WHERE "' || field || '" = ' || hpid || ' AND "to" = ' || other || '
        )';

    END LOOP;

END $$;

CREATE OR REPLACE FUNCTION message_control(message text) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE ret text;
BEGIN
    SELECT trim(message) INTO ret;
    IF char_length(ret) = 0 THEN
        RAISE EXCEPTION 'NO_EMPTY_MESSAGE';
    END IF;
    RETURN ret;
END $$;

CREATE OR REPLACE FUNCTION post_control() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.message = message_control(NEW.message);

    IF TG_OP = 'INSERT' THEN -- no flood control on update
        PERFORM flood_control('"posts"', NEW."from", NEW.message);
    END IF;

    PERFORM blacklist_control(NEW."from", NEW."to");

    IF( NEW."to" <> NEW."from" AND
        (SELECT "closed" FROM "profiles" WHERE "counter" = NEW."to") IS TRUE AND 
        NEW."from" NOT IN (SELECT "to" FROM whitelist WHERE "from" = NEW."to")
      )
    THEN
        RAISE EXCEPTION 'CLOSED_PROFILE';
    END IF;


    IF TG_OP = 'UPDATE' THEN -- no pid increment
        SELECT (now() at time zone 'utc') INTO NEW.time;
    ELSE
        SELECT "pid" INTO NEW.pid FROM (
            SELECT COALESCE( (SELECT "pid" + 1 as "pid" FROM "posts"
            WHERE "to" = NEW."to"
            ORDER BY "hpid" DESC
            FETCH FIRST ROW ONLY), 1 ) AS "pid"
        ) AS T1;
    END IF;

    IF NEW."to" <> NEW."from" THEN -- can't write news to others board
        SELECT false INTO NEW.news;
    END IF;

    -- if to = GLOBAL_NEWS set the news filed to true
    IF NEW."to" = (SELECT counter FROM special_users where "role" = 'GLOBAL_NEWS') THEN
        SELECT true INTO NEW.news;
    END IF;
    
    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION post_update() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
  BEGIN
     INSERT INTO posts_revisions(hpid, time, message, rev_no) VALUES(OLD.hpid, OLD.time, OLD.message,
         (SELECT COUNT(hpid) +1 FROM posts_revisions WHERE hpid = OLD.hpid));

     PERFORM hashtag(NEW.message, NEW.hpid, false, NEW.from, NEW.time);
     PERFORM mention(NEW."from", NEW.message, NEW.hpid, false);
     RETURN NULL;
 END $$;

CREATE OR REPLACE FUNCTION strip_tags(message text) RETURNS text
    LANGUAGE plpgsql
    AS $$
    begin
        return regexp_replace(regexp_replace(
          regexp_replace(regexp_replace(
          regexp_replace(regexp_replace(
          regexp_replace(regexp_replace(
          regexp_replace(message,
             '\[url[^\]]*?\](.*)\[/url\]',' ','gi'),
             '\[code=[^\]]+\].+?\[/code\]',' ','gi'),
             '\aa[c=[^\]]+\].+?\[/c\]',' ','gi'),
             '\[video\].+?\[/video\]',' ','gi'),
             '\[yt\].+?\[/yt\]',' ','gi'),
             '\[youtube\].+?\[/youtube\]',' ','gi'),
             '\[music\].+?\[/music\]',' ','gi'),
             '\[img\].+?\[/img\]',' ','gi'),
             '\[twitter\].+?\[/twitter\]',' ','gi');
    end $$;



--
-- Name: user_comment(); Type: FUNCTION; Schema: public; 
--

CREATE OR REPLACE FUNCTION user_comment() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
     PERFORM hashtag(NEW.message, NEW.hpid, false, NEW.from, NEW.time);
     PERFORM mention(NEW."from", NEW.message, NEW.hpid, false);
     -- edit support
     IF TG_OP = 'UPDATE' THEN
         INSERT INTO comments_revisions(hcid, time, message, rev_no)
         VALUES(OLD.hcid, OLD.time, OLD.message, (
             SELECT COUNT(hcid) + 1 FROM comments_revisions WHERE hcid = OLD.hcid
         ));

          --notify only if it's the last comment in the post
         IF OLD.hcid <> (SELECT MAX(hcid) FROM comments WHERE hpid = NEW.hpid) THEN
             RETURN NULL;
         END IF;
     END IF;

     -- if I commented the post, I stop lurking
     DELETE FROM "lurkers" WHERE "hpid" = NEW."hpid" AND "from" = NEW."from";

     WITH no_notify("user") AS (
         -- blacklist
         (
             SELECT "from" FROM "blacklist" WHERE "to" = NEW."from"
                 UNION
             SELECT "to" FROM "blacklist" WHERE "from" = NEW."from"
         )
         UNION -- users that locked the notifications for all the thread
             SELECT "user" FROM "posts_no_notify" WHERE "hpid" = NEW."hpid"
         UNION -- users that locked notifications from me in this thread
             SELECT "to" FROM "comments_no_notify" WHERE "from" = NEW."from" AND "hpid" = NEW."hpid"
         UNION -- users mentioned in this post (already notified, with the mention)
             SELECT "to" FROM "mentions" WHERE "u_hpid" = NEW.hpid AND to_notify IS TRUE
         UNION
             SELECT NEW."from"
     ),
     to_notify("user") AS (
             SELECT DISTINCT "from" FROM "comments" WHERE "hpid" = NEW."hpid"
         UNION
             SELECT "from" FROM "lurkers" WHERE "hpid" = NEW."hpid"
         UNION
             SELECT "from" FROM "posts" WHERE "hpid" = NEW."hpid"
         UNION
             SELECT "to" FROM "posts" WHERE "hpid" = NEW."hpid"
     ),
     real_notify("user") AS (
         -- avoid to add rows with the same primary key
         SELECT "user" FROM (
             SELECT "user" FROM to_notify
                 EXCEPT
             (
                 SELECT "user" FROM no_notify
              UNION
                 SELECT "to" AS "user" FROM "comments_notify" WHERE "hpid" = NEW."hpid"
             )
         ) AS T1
     )

     INSERT INTO "comments_notify"("from","to","hpid","time") (
         SELECT NEW."from", "user", NEW."hpid", NEW."time" FROM real_notify
     );

     RETURN NULL;
 END $$;

CREATE OR REPLACE FUNCTION user_comment_edit_control() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.editable IS FALSE THEN
        RAISE EXCEPTION 'NOT_EDITABLE';
    END IF;

    -- update time
    SELECT (now() at time zone 'utc') INTO NEW.time;

    NEW.message = message_control(NEW.message);
    PERFORM flood_control('"comments"', NEW."from", NEW.message);
    PERFORM blacklist_control(NEW."from", NEW."to");

    RETURN NEW;
END $$;

CREATE OR REPLACE FUNCTION user_interactions(me bigint, other bigint) RETURNS SETOF record
    LANGUAGE plpgsql
    AS $$
DECLARE tbl text;
        ret record;
        query text;
begin
    FOR tbl IN (SELECT unnest(array['blacklist', 'comment_thumbs', 'comments', 'followers', 'lurkers', 'mentions', 'pms', 'posts', 'whitelist'])) LOOP
        query := interactions_query_builder(tbl, me, other, false);
        FOR ret IN EXECUTE query LOOP
            RETURN NEXT ret;
        END LOOP;
    END LOOP;
   RETURN;
END $$;

ALTER TABLE ban ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE ban ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE blacklist ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE blacklist ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE bookmarks ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE bookmarks ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE comment_thumbs ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE comment_thumbs ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE comments ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE comments ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE comments_no_notify ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE comments_no_notify ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE comments_notify ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE comments_notify ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE comments_revisions ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE comments_revisions ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE deleted_users ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE deleted_users ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups ALTER COLUMN creation_time TYPE timestamp without time zone;
ALTER TABLE groups ALTER COLUMN creation_time SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_bookmarks ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_bookmarks ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_comment_thumbs ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_comment_thumbs ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_comments ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_comments ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_comments_no_notify ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_comments_no_notify ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_comments_notify ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_comments_notify ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_comments_revisions ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_comments_revisions ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_followers ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_followers ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_lurkers ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_lurkers ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_members ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_members ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_notify ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_notify ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_owners ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_owners ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_posts ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_posts ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_posts_no_notify ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_posts_no_notify ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_posts_revisions ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_posts_revisions ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE groups_thumbs ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE groups_thumbs ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE guests ALTER COLUMN last TYPE timestamp without time zone;
ALTER TABLE guests ALTER COLUMN last SET DEFAULT now() at time zone 'utc';

ALTER TABLE lurkers ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE lurkers ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE mentions ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE mentions ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE posts ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE posts ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE posts_classification ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE posts_classification ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE posts_no_notify ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE posts_no_notify ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE posts_notify ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE posts_notify ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE posts_revisions ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE posts_revisions ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE profiles ALTER COLUMN pushregtime TYPE timestamp without time zone;
ALTER TABLE profiles ALTER COLUMN pushregtime SET DEFAULT now() at time zone 'utc';

ALTER TABLE reset_requests ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE reset_requests ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE searches ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE searches ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

ALTER TABLE users ALTER COLUMN last TYPE timestamp without time zone;
ALTER TABLE users ALTER COLUMN last SET DEFAULT now() at time zone 'utc';

ALTER TABLE whitelist ALTER COLUMN "time" TYPE timestamp without time zone;
ALTER TABLE whitelist ALTER COLUMN "time" SET DEFAULT now() at time zone 'utc';

create view messages as
select "hpid","from","to","pid","message","time","news","lang","closed", 0 as type from groups_posts
union all
select "hpid","from","to","pid","message","time","news","lang","closed", 1 as type from posts;

update profiles set template = '0', mobile_template = '0';

drop table if exists interests cascade;

create table interests(
    id bigserial primary key not null,
    "from" bigint not null references users(counter) on delete cascade,
    value varchar(90) not null,
    time timestamp without time zone default (now() at time zone 'utc')
);

create unique index "unique_intersest_from_value" on interests("from", LOWER("value"));

insert into interests("from", value)
select distinct a,b from (select counter as a, unnest(arr) as b from (select counter, regexp_split_to_array(interests, E'\\s*\\n\\s*') as arr from profiles) t where arr <> '{""}') x where length(b) <= 90;

alter table profiles drop column interests;

-- dateformat is only for the date, not for the time
alter table profiles alter column dateformat set default 'd/m/Y';
update profiles set dateformat = 'd/m/Y';

-- TODO: https://news.ycombinator.com/item?id=9512912
-- https://blog.lateral.io/2015/05/full-text-search-in-milliseconds-with-postgresql/
/*
alter table posts add column tsv tsvector;
alter table posts_comments add column tsv tsvector;
alter table groups_posts add column tsv tsvector;
alter table groups_posts_comments add column tsv tsvector;
 */

commit;
