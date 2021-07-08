--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 13.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.UTF-8';


ALTER DATABASE postgres OWNER TO postgres;

\connect postgres

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id text,
    first_name text,
    last_name text,
    username text,
    password text,
    secret text
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public."user" VALUES ('210708212909657cb25f', 'fahmy', 'abida', 'fahmyabida', '410b4a22e2ee767d883390ee719d5ea6dc6d33c1177a73b76c7de8166e20e69a', '1cdee6dee8');
INSERT INTO public."user" VALUES ('2107082131328312bd20', 'frank', 'sinatra', 'userfrank', '57b04cbda9b5b45aee9c8a4c036529e9dc39333b58a06b1f150cb93d51794577', '5e2d4f52d0');


--
-- PostgreSQL database dump complete
--

