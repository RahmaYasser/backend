--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: contenttype; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contenttype (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    name character varying(50) NOT NULL
);


ALTER TABLE public.contenttype OWNER TO postgres;

--
-- Name: contenttype_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.contenttype_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.contenttype_id_seq OWNER TO postgres;

--
-- Name: contenttype_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.contenttype_id_seq OWNED BY public.contenttype.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: tamiat_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tamiat_permissions (
    id integer NOT NULL,
    name character varying(60)
);


ALTER TABLE public.tamiat_permissions OWNER TO postgres;

--
-- Name: tamiat_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tamiat_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tamiat_permissions_id_seq OWNER TO postgres;

--
-- Name: tamiat_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tamiat_permissions_id_seq OWNED BY public.tamiat_permissions.id;


--
-- Name: tamiat_role_permission; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tamiat_role_permission (
    id integer NOT NULL,
    role_id integer,
    permission_id integer
);


ALTER TABLE public.tamiat_role_permission OWNER TO postgres;

--
-- Name: tamiat_role_permission_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tamiat_role_permission_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tamiat_role_permission_id_seq OWNER TO postgres;

--
-- Name: tamiat_role_permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tamiat_role_permission_id_seq OWNED BY public.tamiat_role_permission.id;


--
-- Name: tamiat_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tamiat_roles (
    id integer NOT NULL,
    name character varying(20)
);


ALTER TABLE public.tamiat_roles OWNER TO postgres;

--
-- Name: tamiat_roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tamiat_roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tamiat_roles_id_seq OWNER TO postgres;

--
-- Name: tamiat_roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tamiat_roles_id_seq OWNED BY public.tamiat_roles.id;


--
-- Name: tamiat_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tamiat_users (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    name character varying(20) NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    role_id integer
);


ALTER TABLE public.tamiat_users OWNER TO postgres;

--
-- Name: tamiat_users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tamiat_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tamiat_users_id_seq OWNER TO postgres;

--
-- Name: tamiat_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tamiat_users_id_seq OWNED BY public.tamiat_users.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    email text NOT NULL,
    password text NOT NULL,
    otp text,
    email_verified boolean DEFAULT false
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: contenttype id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contenttype ALTER COLUMN id SET DEFAULT nextval('public.contenttype_id_seq'::regclass);


--
-- Name: tamiat_permissions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_permissions ALTER COLUMN id SET DEFAULT nextval('public.tamiat_permissions_id_seq'::regclass);


--
-- Name: tamiat_role_permission id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_role_permission ALTER COLUMN id SET DEFAULT nextval('public.tamiat_role_permission_id_seq'::regclass);


--
-- Name: tamiat_roles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_roles ALTER COLUMN id SET DEFAULT nextval('public.tamiat_roles_id_seq'::regclass);


--
-- Name: tamiat_users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_users ALTER COLUMN id SET DEFAULT nextval('public.tamiat_users_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: contenttype contenttype_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contenttype
    ADD CONSTRAINT contenttype_pkey PRIMARY KEY (id);


--
-- Name: tamiat_permissions tamiat_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_permissions
    ADD CONSTRAINT tamiat_permissions_pkey PRIMARY KEY (id);


--
-- Name: tamiat_role_permission tamiat_role_permission_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_role_permission
    ADD CONSTRAINT tamiat_role_permission_pkey PRIMARY KEY (id);


--
-- Name: tamiat_roles tamiat_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_roles
    ADD CONSTRAINT tamiat_roles_pkey PRIMARY KEY (id);


--
-- Name: tamiat_users tamiat_users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_users
    ADD CONSTRAINT tamiat_users_email_key UNIQUE (email);


--
-- Name: tamiat_users tamiat_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_users
    ADD CONSTRAINT tamiat_users_pkey PRIMARY KEY (id);


--
-- Name: contenttype unique_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contenttype
    ADD CONSTRAINT unique_name UNIQUE (name);


--
-- Name: tamiat_permissions unique_permissions_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_permissions
    ADD CONSTRAINT unique_permissions_name UNIQUE (name);


--
-- Name: tamiat_roles unique_roles_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_roles
    ADD CONSTRAINT unique_roles_name UNIQUE (name);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: tamiat_role_permission tamiat_role_permission_permission_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_role_permission
    ADD CONSTRAINT tamiat_role_permission_permission_id_fkey FOREIGN KEY (permission_id) REFERENCES public.tamiat_permissions(id);


--
-- Name: tamiat_role_permission tamiat_role_permission_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_role_permission
    ADD CONSTRAINT tamiat_role_permission_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.tamiat_roles(id);


--
-- Name: tamiat_users tamiat_users_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tamiat_users
    ADD CONSTRAINT tamiat_users_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.tamiat_roles(id);


--
-- PostgreSQL database dump complete
--

