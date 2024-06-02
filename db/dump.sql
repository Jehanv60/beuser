--
-- PostgreSQL database dump
--

-- Dumped from database version 12.18 (Ubuntu 12.18-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.18 (Ubuntu 12.18-0ubuntu0.20.04.1)

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
-- Name: pengguna; Type: TABLE; Schema: public; Owner: han
--

CREATE TABLE public.pengguna (
    id integer NOT NULL,
    pengguna character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(100) NOT NULL
);


ALTER TABLE public.pengguna OWNER TO han;

--
-- Name: pengguna_id_seq; Type: SEQUENCE; Schema: public; Owner: han
--

CREATE SEQUENCE public.pengguna_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pengguna_id_seq OWNER TO han;

--
-- Name: pengguna_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: han
--

ALTER SEQUENCE public.pengguna_id_seq OWNED BY public.pengguna.id;


--
-- Name: pengguna id; Type: DEFAULT; Schema: public; Owner: han
--

ALTER TABLE ONLY public.pengguna ALTER COLUMN id SET DEFAULT nextval('public.pengguna_id_seq'::regclass);


--
-- Data for Name: pengguna; Type: TABLE DATA; Schema: public; Owner: han
--

COPY public.pengguna (id, pengguna, email, password) FROM stdin;
2	hangg1	gege1@gmail.com	$2a$10$rwOCmYHHZYsgsTbVw8773.Wtx2/Zmdm3b.M0D61jygjvDlIbVNtMi
3	hangg3	gege3@gmail.com	$2a$10$XUYuSOYhUZUoPo/04sGkF.W1umM5/vIjUYCQNRBS8EDV6IJpifIpC
1	hangg	gege@gmail.com	$2a$10$s6BW/y7E/wFBrRIS3BGbHe37FHoSm6bWE2cuvRcmQKtBmg0SNJWFK
\.


--
-- Name: pengguna_id_seq; Type: SEQUENCE SET; Schema: public; Owner: han
--

SELECT pg_catalog.setval('public.pengguna_id_seq', 4, true);


--
-- Name: pengguna pengguna_pkey; Type: CONSTRAINT; Schema: public; Owner: han
--

ALTER TABLE ONLY public.pengguna
    ADD CONSTRAINT pengguna_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

