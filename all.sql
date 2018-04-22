--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.5
-- Dumped by pg_dump version 9.6.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: tables; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA tables;


ALTER SCHEMA tables OWNER TO postgres;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = tables, pg_catalog;

--
-- Name: check_items_quantity(); Type: FUNCTION; Schema: tables; Owner: postgres
--

CREATE FUNCTION check_items_quantity() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
declare
	quantity integer;
begin
    if TG_OP = 'INSERT'  then
	    SELECT tables.order_product.quantity FROM tables.order_product WHERE id_product = new.id_product and id_order=new.id_order and id_size=new.id_size into quantity;
	    if (quantity is not null) then 
            update tables.order_product 
            set quantity = (select tables.order_product.quantity from tables.order_product where id_product = new.id_product and id_order=new.id_order and id_size=new.id_size) + 1 
            where id_product = new.id_product and id_order=new.id_order and id_size=new.id_size;
        elsif  (quantity is null) then 
            return new;
        end if;	
    end if;     
    return null;
end;   
$$;


ALTER FUNCTION tables.check_items_quantity() OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: catalog; Type: TABLE; Schema: tables; Owner: postgres
--

CREATE TABLE catalog (
    id integer NOT NULL,
    title text NOT NULL,
    id_parent integer
);


ALTER TABLE catalog OWNER TO postgres;

--
-- Name: order_product; Type: TABLE; Schema: tables; Owner: postgres
--

CREATE TABLE order_product (
    id integer NOT NULL,
    id_product integer NOT NULL,
    id_order integer NOT NULL,
    quantity integer DEFAULT 1 NOT NULL,
    id_size integer NOT NULL
);


ALTER TABLE order_product OWNER TO postgres;

--
-- Name: order_product_id_seq; Type: SEQUENCE; Schema: tables; Owner: postgres
--

CREATE SEQUENCE order_product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_product_id_seq OWNER TO postgres;

--
-- Name: order_product_id_seq; Type: SEQUENCE OWNED BY; Schema: tables; Owner: postgres
--

ALTER SEQUENCE order_product_id_seq OWNED BY order_product.id;


--
-- Name: orders; Type: TABLE; Schema: tables; Owner: postgres
--

CREATE TABLE orders (
    id integer NOT NULL,
    number text NOT NULL,
    cost integer DEFAULT 0 NOT NULL,
    id_user bigint NOT NULL,
    status text DEFAULT 'in processing'::text NOT NULL
);


ALTER TABLE orders OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: tables; Owner: postgres
--

CREATE SEQUENCE orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE orders_id_seq OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: tables; Owner: postgres
--

ALTER SEQUENCE orders_id_seq OWNED BY orders.id;


--
-- Name: product_sizes; Type: TABLE; Schema: tables; Owner: postgres
--

CREATE TABLE product_sizes (
    id_product integer NOT NULL,
    id_sizes integer NOT NULL,
    quantity integer NOT NULL
);


ALTER TABLE product_sizes OWNER TO postgres;

--
-- Name: products; Type: TABLE; Schema: tables; Owner: postgres
--

CREATE TABLE products (
    id integer NOT NULL,
    title text NOT NULL,
    price integer NOT NULL,
    color text NOT NULL,
    id_category integer NOT NULL,
    description text NOT NULL,
    photo text NOT NULL
);


ALTER TABLE products OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: tables; Owner: postgres
--

CREATE SEQUENCE products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE products_id_seq OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: tables; Owner: postgres
--

ALTER SEQUENCE products_id_seq OWNED BY products.id;


--
-- Name: reviews; Type: TABLE; Schema: tables; Owner: postgres
--

CREATE TABLE reviews (
    id integer NOT NULL,
    id_product integer NOT NULL,
    id_user bigint NOT NULL,
    date text NOT NULL,
    description text NOT NULL
);


ALTER TABLE reviews OWNER TO postgres;

--
-- Name: sizes; Type: TABLE; Schema: tables; Owner: postgres
--

CREATE TABLE sizes (
    id integer NOT NULL,
    title text NOT NULL
);


ALTER TABLE sizes OWNER TO postgres;

--
-- Name: sizes_id_seq; Type: SEQUENCE; Schema: tables; Owner: postgres
--

CREATE SEQUENCE sizes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sizes_id_seq OWNER TO postgres;

--
-- Name: sizes_id_seq; Type: SEQUENCE OWNED BY; Schema: tables; Owner: postgres
--

ALTER SEQUENCE sizes_id_seq OWNED BY sizes.id;


--
-- Name: users; Type: TABLE; Schema: tables; Owner: postgres
--

CREATE TABLE users (
    id bigint NOT NULL,
    username text DEFAULT 'none'::text,
    phone text DEFAULT 'none'::text,
    fullname text DEFAULT 'none'::text,
    address text DEFAULT 'none'::text,
    registration_completed boolean DEFAULT false,
    id_current integer DEFAULT 3,
    current_offset integer DEFAULT 0
);


ALTER TABLE users OWNER TO postgres;

--
-- Name: order_product id; Type: DEFAULT; Schema: tables; Owner: postgres
--

ALTER TABLE ONLY order_product ALTER COLUMN id SET DEFAULT nextval('order_product_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: tables; Owner: postgres
--

ALTER TABLE ONLY orders ALTER COLUMN id SET DEFAULT nextval('orders_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: tables; Owner: postgres
--

ALTER TABLE ONLY products ALTER COLUMN id SET DEFAULT nextval('products_id_seq'::regclass);


--
-- Name: sizes id; Type: DEFAULT; Schema: tables; Owner: postgres
--

ALTER TABLE ONLY sizes ALTER COLUMN id SET DEFAULT nextval('sizes_id_seq'::regclass);


--
-- Data for Name: catalog; Type: TABLE DATA; Schema: tables; Owner: postgres
--

COPY catalog (id, title, id_parent) FROM stdin;
0	Корень	\N
1	Одежда	0
2	Обувь	0
3	Женская одежда	1
4	Мужская одежда	1
5	Женская обувь	2
6	Мужская обувь	2
7	Верхняя одежда	3
8	Джемперы и толстовки	3
9	Блузки и рубашки	3
10	Брюки и джинсы	3
11	Комбинезоны	3
12	Платья	3
13	Футболки и майки	3
14	Жилеты	3
15	Юбки	3
16	Куртки	7
17	Пальто	7
18	Шубы	7
19	Дубленки	7
20	Плащи	7
21	Джемперы	8
22	Толстовки	8
23	Блузки	9
24	Рубашки	9
25	Брюки	10
26	Джинсы	10
27	Футболки	13
28	Майки	13
29	Верхняя одежда	4
30	Джемперы и свитеры	4
31	Брюки и джинсы	4
32	Рубашки	4
33	Футболки и майки	4
34	Пиджаки и жакеты	4
35	Толстовки	4
36	Куртки	29
37	Пальто	29
38	Бомберы	29
39	Парки	29
40	Джемперы	30
41	Свитеры	30
42	Брюки	31
43	Джинсы	31
44	Футболки	33
45	Майки	33
46	Пиджаки	34
47	Жакеты	34
48	Балетки	5
49	Ботильоны	5
50	Ботинки	5
51	Кеды	5
52	Сапоги	5
53	Туфли	5
54	Тапочки	5
55	Шлепанцы	5
56	Ботинки	6
57	Кеды	6
58	Полуботинки	6
59	Полусапоги	6
60	Сапоги	6
61	Спортивная обувь	6
62	Тапочки	6
63	Туфли	6
64	trty54y	3
65	qwertyu	3
100	ЛЯЛЯЛЯЛЯ	13
\.


--
-- Data for Name: order_product; Type: TABLE DATA; Schema: tables; Owner: postgres
--

COPY order_product (id, id_product, id_order, quantity, id_size) FROM stdin;
28	5	3	7	1
15	4	4	1	3
14	5	4	1	1
30	3	3	1	2
31	3	3	1	5
\.


--
-- Name: order_product_id_seq; Type: SEQUENCE SET; Schema: tables; Owner: postgres
--

SELECT pg_catalog.setval('order_product_id_seq', 43, true);


--
-- Data for Name: orders; Type: TABLE DATA; Schema: tables; Owner: postgres
--

COPY orders (id, number, cost, id_user, status) FROM stdin;
3	BpLnfgDsc3	0	364794408	in processing
4	BpLnfgDsc3	0	294176487	in processing
5	BpLnfgDsc3	0	326419	in processing
\.


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: tables; Owner: postgres
--

SELECT pg_catalog.setval('orders_id_seq', 5, true);


--
-- Data for Name: product_sizes; Type: TABLE DATA; Schema: tables; Owner: postgres
--

COPY product_sizes (id_product, id_sizes, quantity) FROM stdin;
1	1	3
1	2	3
1	3	3
2	2	3
2	1	3
3	2	3
3	1	3
3	5	3
4	3	3
5	1	3
5	1	3
6	2	3
50	1	3
50	5	3
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: tables; Owner: postgres
--

COPY products (id, title, price, color, id_category, description, photo) FROM stdin;
1	Юбка в складку	1799	Белый с цветочным принтом	15	Выполнена из легкого хлопкового полотна, линия талии на среднем уровне, боковые карманы, длина миди, состав:98% хлопок, 2% эластан	AgADAgADKqkxG5FEUUiKa_FTXrOSIr_qAw4ABLBAVhiH1Wo7VWwDAAEC
2	Юбка А-силуэта	2699	Тёмно-бардовый	15	Выполнена из лиоцелла, высокое положение линии талии, декоративный текстильный пояс, боковые карманы, средняя длина, состав: 100% лиоцелл	AgADAgADLKkxG5FEUUgLgeLRiat5Qxabmg4ABBx_iDGrvr5oGY0BAAEC
3	Юбка из сетки	1399	Персиковый	15	Выполнена из сетки на подкладке, линия талии на среднем уровне, эластичный пояс, длина миди, состав: 100% полиэстер	AgADAgADLakxG5FEUUiOvLiK0SQtuyIKnA4ABJtl5P0NwPatI48BAAEC
4	Юбка-карандаш	1999	Красный	15	Выполнена из эластичного хлопка, высокая талия, декоративный ремень, боковые карманы, средняя длина, состав:97% хлопок, 3% эластан	AgADAgADLqkxG5FEUUg7u0Poj9ajpgb5Mg4ABFK-q_spL2abRDEEAAEC
5	Юбка из жаккарда	1999	Хаки	15	Высокая линия талии, этнический жаккард, юбка-мини, состав: 99% хлопок, 1% металлизированное волокно	AgADAgADL6kxG5FEUUhcxq33FZ798hzuAw4ABFH0wqE8ZRhbQm4DAAEC
6	Юбка в из искусственной кожи	1599	Черный	15	Выполнена из искусственной кожи, нормальное положение линии талии, длина миди, серебряная молния-застежка спереди, состав: искусственная кожа	AgADAgADMKkxG5FEUUhjmvGqJR5Y9LbqAw4ABAL2gQF5H9WPV2oDAAEC
7	Жилет с вышивкой	2499	Джинсовый	14	Выполнен из мягкого денимного полотна, заострённый воротник, вышивка на спине в виде дракона, свободный удлиненный силуэт, состав: 100% хлопок	AgADAgADKKkxG5FEUUjX8wKnxYmyUJQInA4ABCb6TNjwujDG-I4BAAEC
8	Удлиненный жилет	2799	Черный	14	Выполнен из поливискозы, без застёжки, боковые карманы, пояс из основной ткани, состав: 65% полиэстер, 35% вискоза	AgADAgADKakxG5FEUUjn5nFoFRlh_IkQMw4ABJge5HhbnRAnnjYEAAEC
9	Расклешённое платье	1399	Светло-розовый	12	Выполнено из эластичного полиэстрового полотна, круглый вырез горловины, ракава длиной 3/4, составЖ 88% полиэстер, 12% эластан	AgADAgADbqgxGymRYEhLLIbp8WPz1gnwAw4ABBFC3Et15GB9ZHUDAAEC
10	Платье из жаккарда	1599	Молочный	12	Выполнено из жаккардного полотна, приталенный силуэт платья, V-образный вырез горловины, 3/4 рукав, длина короткая, состав: 97% полиэстер, 3% эластан	AgADAgADb6gxGymRYEhuTSgIv8PnkS8DnA4ABHl1Ywan-k6WiJgBAAEC
11	Платье в цветочный принт	1799	Черный	12	Выполнено из трикотажа креп-структуры, приталенное платье с пышной юбкой, длина короткая, состав 97% полиэстер, 3% эластан	AgADAgADcKgxGymRYEimJ89xrO3UGTn2Aw4ABOxJr1XmowABy55yAwABAg
12	Платье с оборкой на плечах	1799	Черный	12	Выполнено из мягкого полиэстрового полотна, V-образный вырез горловины, бант сзади, рюша спереди и сзади, состав: 100% полиэстер	AgADAgADcagxGymRYEiqE6NPD2cda5oInA4ABEqsexKR3omVMZUBAAEC
13	Платье-толстовка	1599	Светло-серый	12	Выполнено из мягкого трикотажа, прямое платье, свободный силуэт, объёмные рукава, длина короткая, капюшон, состав: 85% хлопок, 15% полиэстер	AgADAgADcqgxGymRYEjwvAo-dn7A3vwIMw4ABKUAAWFNeT2CAoU-BAABAg
14	Платье из искусственной замши	2499	Светло-красный	12	Выполнено из искусственной замши,круглый вырез горловины, состав: 85% полиэстер, 15% эластан	AgADAgADc6gxGymRYEiVUy5QBpffprUbMw4ABKcCmi73d_jozD0EAAEC
15	Платье-рубашка	1999	Синий	12	Выполнено из лиоцелла, платье-рубашка, прямое по силуэту, застёжка на пуговицы, съёмный пояс, состав: 100% лиоцелл	AgADAgADdKgxGymRYEjSz-lnXkFuPuD9Mg4ABKfGJQaWg5fwUT0EAAEC
16	Прямое платье	1999	Черный	12	Выполнено из структурного полиэстера, прямое платье, короткие рукава, круглый вырез горловины, средняя длина, состав: 96% полиэстер, 4% эластан	AgADAgADdagxGymRYEgUaQABjlvSHEDF6AMOAATZ14PWy15yfKt0AwABAg
17	Трикотажное платье	1799	Тёмно-синий	12	Выполнено из прямого структурного трикотажа, прямой силуэт, короткие рукава, средняя длина, состав: 98% полиэстер, 2% эластан	AgADAgADdqgxGymRYEjAayMfcP9da7gHMw4ABGWRH-QQzo12CDoEAAEC
18	Платье-карандаш	1999	Тёмно-синий	12	Выполнено из плотного интерлока с вискозой, платье прилигающее, короткие рукава, средняя длина, состав: 67% вискоза, 28% нейлон, 5% эластан	AgADAgADd6gxGymRYEhAoBlpk2Ywas0aMw4ABMEG533HFYi4OjoEAAEC
19	Трикотажное платье	1599	Черно-синий	12	Выполнено из плотного структурного трикотажа, прямой силуэт, круглый вырез горловины, короткие рукава, состав: 98% полиэстер, 2% эластан	AgADAgADeKgxGymRYEhmckq_52W2cur0mw4ABBAqRwABdYIjB22RAQABAg
20	Платье из замши	1999	Персиковый	12	Выполнено из искусственной мягкой замши, V-образный вырез горловины, боковые карманы, короткие рукава, длина миди, состав: 100% полиэстер	AgADAgADeagxGymRYEjUkTw4MOmeGPgRMw4ABOCXYFM3PjByI0IEAAEC
43	Пальто из норки	79900	Чёрный	18	Красивая шуба классического цвета, отрезное по талии, по бокам прорезные карманы	AgADAgADT6kxG1ELmUhpQmj37VO5QzMWnA4ABNrFaZVb_RJ7XdQBAAEC
21	Платье в микропринт	1799	Светло-серый	12	Выполнено из жаккардного полотна, короткие рукава, круглый вырез горловины, средняя длина, состав: 75% полиэстер, 25% вискоза	AgADAgADeqgxGymRYEgQ1Ydw5dXLjNDvAw4ABD_7O7khcuexwnIDAAEC
22	Платье-толстовка	1199	Светло-серый	12	Выполнено из мягкого трикотажа, круглый вырез горловины, спущенные плечи, асиметричная линия низа, состав: 95% хлопок, 5% вискоза	AgADAgADe6gxGymRYEjNZNHlEslo2kkJMw4ABCQMhBZTJiGL7TwEAAEC
23	Платье из шифона	2999	Тёмно-синий	12	Выполнено из мягкого полиэстрового шифона, круглый вырез, рукава 3/4, низ платья в складку,средняя длина, состав: 100% полиэстер	AgADAgADfKgxGymRYEhrW6U0xGs8V1r7Mg4ABNyHUE6b9lYNZj4EAAEC
24	Платье-труба	1999	Чёрный	12	Выполнено из плотного вискозного полотна, приталенный силуэт, молния на спинке, круглый вырез горловины, состав: 68% вискоза, 28% нейлон, 4% эластан	AgADAgADfagxGymRYEjrIjJAyyGIWAzuAw4ABOVc3zvX24-orHMDAAEC
25	Платье в цветочный принт	1999	Индиго	12	Выполнено из крепового полиэстрового полотна, корсетные детали по бокам, молния на спинке, длинные рукава, длина средняя, состав: 100% полиэстер	AgADAgADfqgxGymRYEhuf66NomfXZbPqAw4ABDSqoJrfWclPpnQDAAEC
26	Комбинезон из денима	2499	Светло-синий	11	Выполнен из джинсовой ткани, карманы на груди и сзади, боковые карманы, эффект порванности, состав: 74% хлопок, 20% полиэстер, 6% вискоза	AgADAgADbKgxGymRYEixXJ0h22seLwr0Aw4ABFA5VmVB6SURf3IDAAEC
27	Комбинезон	4599	Чёрный	11	Состав: 96% вискоза, 4% эластан	AgADAgADbagxGymRYEjzlgABE7xkDJ-Hm0MOAARAUbkdzkqviv-WAQABAg
28	Кожаная куртка	2999	Чёрный	16	Куртка из экокожи с перфорацией, без воротника, застёжка на молнию, боковые карманы, состав: 100% полиуретан	AgADAgADFKkxG2RKmEj1p5tXEQsd2FD1Aw4ABEccKXxrWGWGyrEDAAEC
29	Парка	3999	Тёмно-синий	16	Аутдорная куртка из технологичной ткани, с полосатым джерси в капюшоне, с фурнитурой контрастного цвета, состав: 100% полиэстер	AgADAgADQakxG1ELmUjdrLx6uybiyGYOnA4ABMKOXyUNwqva3tIBAAEC
30	Парка с отделкой	3999	Светло-персиковый	16	Длинная легкая парка, отделка контрастного цвета, застёжка на молнии, состав: 100% хлопок	AgADAgADQqkxG1ELmUhH3o0AAQ4nKmqeDJwOAASVKfrHq4DSpF_SAQABAg
31	Кожаная куртка	3999	Светло-бежевый	16	Куртка из экокожи с центральной молнией, без воротника, с рюшей по низу и по плечам, состав: 100% полиуретан	AgADAgADQ6kxG1ELmUhQPaSvF17cA4DyAw4ABElh8ZU2z288I7ADAAEC
32	Сатиновый бомбер	3499	Светло-персиковый	16	Бомбер с утеплителем, с ромбовидной стёжкой, воротник гольф, карман на рукаве, состав: 100% полиэстер	AgADAgADRKkxG1ELmUgZ47LqktBXTvy3mg4ABCXmH_c9xmnKcdIBAAEC
33	Приталенная куртка-блейзер	3499	Чёрно-синий	16	Ультра-легкий стёганый блейзер, приталенный силуэт, застёжка на кнопках, состав: 100% нейлон	AgADAgADRakxG1ELmUgKSxH5-_lxs7_3Mg4ABLrJqWSTBljNh3YEAAEC
34	Куртка на кнопках	2999	Серый	16	Ультра-легкая стёганая куртка без воротника, силуэт оверсайз, застёжка на кнопках, состав: 100% нейлон	AgADAgADRqkxG1ELmUj4QY77q-KvyuP8Mg4ABPpxJ-TGTEUGT4AEAAEC
35	Трикотажное пальто	3999	Чёрно-синий	17	Неопреновое пальто из ткани с контрасной изнанкой с капюшоном, без подкладки, состав: 66% полиэстер, 29% вискоза, 5% эластан	AgADAgADR6kxG1ELmUiBY6bf1YKmMcz8Mg4ABPpw0CsbXIgNGnoEAAEC
36	Полупальто с воротником-стойкой	4999	Красный	17	Формальное пальто, прямой силуэт, рукав-реглан, воротник-стойка, состав: 68% полиэстер, 32% вискоза	AgADAgADSKkxG1ELmUhBJUxr5GfPwTsOnA4ABFzXm4RPaZwP79IBAAEC
37	Приталенное пальто с капюшоном	3999	Серый	17	Из трикотажной ткани, приталенный силуэт, объёмный капюшон, состав: 52% полиэстер, 31% хлопок, 17% вискоза	AgADAgADSakxG1ELmUhmdAPUH0Bq7-8GMw4ABMrt57ZETTal03cEAAEC
38	Полупальто на молнии	4999	Чёрно-синий	17	Без воротника, трапециевидный силуэт, застёжка на молнии, боковые карманы, состав: 64% полиэстер, 36% вискоза	AgADAgADSqkxG1ELmUjvpZIJejKGORL4Mg4ABO7K7WayaV0Ly3MEAAEC
39	Шерстяное пальто	4999	Светло-розовый	17	Пальто с английским воротником, силуэт оверсайз, ткань с содержанием шерсти, состав: 65% полиэстер, 35% шерсть	AgADAgADS6kxG1ELmUikE0uc_ceOskkHnA4ABPlKNnZVHr_7ddABAAEC
40	Пальто с воротником "гольф"	3999	Серый	17	Ультра-лёгкое стёганое пальто, прямой силуэт, боковые наклдные, застёжка на молнии, состав: 100% полиэстер	AgADAgADTKkxG1ELmUgf4l_K5z0KxgEanA4ABA104_I25X-RhM8BAAEC
41	Приталенное пальто без воротника	3999	Розовый	17	Ультра-легкое пальто без воротника, кулиска по талии, принтовая подкладка, состав: 100% нейлон	AgADAgADTakxG1ELmUhtowjJ82sX4wwYnA4ABHIJLqzwUIeMP9EBAAEC
42	Пальто с капюшоном	3999	Голубой	17	Ультра-лёгкое утепленное пальто с рукавами и капюшоном из курточной ткани, приталенный силуэт, состав: 100% нейлон 	AgADAgADTqkxG1ELmUiCYWk-kQqKa9wXnA4ABIBKqWgyR--vhdMBAAEC
110	Топ	1999	Голубой	28	Состав: 70% хлопок, 30% полиэстер	AgADAgADjakxG1ELmUi6nztdadukkkUKMw4ABE3ANyEyAlWmtH0EAAEC
44	Пальто из норки	99900	Молочный	18	Роскошное норковое пальто, уникальный цвет, расклешенная нижняя часть, капюшон с отворотом	AgADAgADUKkxG1ELmUg4mKlyJCZDjrvoAw4ABP697mGUSXixpK0DAAEC
45	Норковая шуба	69000	Тёмно-коричневый	18	Силуэт-полутрапеция, технология пошива в роспуск, двойной капюшон, прорезные карманы	AgADAgADUakxG1ELmUjEkkDnuFDMaY8PnA4ABBMuCejg4pg5Mc8BAAEC
46	Дублёнка	4599	Светло-розовый	19	Выполнена из плотного текстиля, приталенный крой, отложной воротникс меховой отделкой, состав: 98% полиэстер, 2% эластан	AgADAgADiqgxG2RKoEjzKu0MRO9-J5cUMw4ABE1yOyNdM0ulgYEEAAEC
47	Дублёнка	6899	Чёрный	19	Выполнена из гладкой искусственной кожи, отложной воротник, застежка на асимметричную молнию, состав: 81% вискоза, 19% полиэстер	AgADAgADUqkxG1ELmUjeqBmWXPlJPzgOnA4ABHUwtPCYB7GpitMBAAEC
48	Дублёнка	7199	Молочный	19	Выполнена из искусственной дубленной замши, прямой крой, застёжка на молнию, два кармана, состав: 100% полиэстер	AgADAgADU6kxG1ELmUhUpAABjIrpmdGDCpwOAATmyBIBT_SCAAFz1QEAAQI
49	Дублёнка	3140	Черный	19	Выполнена из искусственной замши и фактурного трикотажа, приталенного кроя, подкладка из искусственного меха, состав: 100% полиэстер	AgADAgADVKkxG1ELmUgq-G9TgXRjQGABMw4ABPQGeKhR5FKqNXUEAAEC
50	Плащ	4499	Молочный	20	Выполнен из плотного хлопкового текстиля, гладкая подкладка, приталенный крой, состав: 98% хлопок, 2% эластан	AgADAgADVakxG1ELmUi7ArM6P8TtrB0WMw4ABCfQGKVtDtoZaH8EAAEC
51	Плащ	5999	Чёрный	20	Выполнен из текстиля, модель прямого кроя, хлястики на рукавах, застёжка на кнопки, состав: 100% полиэстер	AgADAgADVqkxG1ELmUjtgJPeuHjRAq8FMw4ABDE9-fN_WmdwYHsEAAEC
52	Плащ	4999	Белый	20	Легкий текстильный плащ, застёгивается на кнопки, два кармана, пояс, состав: 50% лиоцелл, 35% лён, 15% полиэстер	AgADAgADV6kxG1ELmUgXP0mc5QABRSPlC5wOAAToxiXEE-hyVt_PAQABAg
53	Плащ	3699	Красный	20	Выполнен из плотного текстиля, приталенный крой с поясом, съёмный капюшон, состав: 60% полиэстер, 40% хлопок	AgADAgADWKkxG1ELmUja5Efh7j4RUoLymw4ABA6Qqkhjn8aNS9EBAAEC
54	Джемпер с открытыми плечами	999	Голубой	21	Выполнен из меланжевого полотна с большим содержанием вискозы, с открытыми плечами, рукава 3/4, состав: 71% вискоза, 29% полиэстер	AgADAgADWqkxG1ELmUhXtyJBPGak6TUGnA4ABKI63M4D8Fq6tdcBAAEC
55	Джемпер с коротким рукавом	1599	Молочный	21	Выполнен из плотного интерлока с вискозой, вырез горловины "лодочка", короткие рукава, состав: 57% вискоза, 43% полиэстер	AgADAgADW6kxG1ELmUgnzAuGdKjZGpP0Aw4ABDmUdxz9-rw1ULEDAAEC
56	Джемпер в полоску с воланами	999	Тёмно-синий	21	Выполнен из вискозного полотна в полоску, круглый вырез горловины, рукава 3/4, состав: 95% вискоза, 5% эластан	AgADAgADjKgxG2RKoEg8BrCKlTfxAoENMw4ABLUgPdiw1uDKl30EAAEC
57	Джемпер в полоску	1599	Тёмно-синий	21	Выполнен из структурного полотна в полоску, короткие рукава, боковые ленты, круглый вырез горловины, состав: 78% полиэстер, 22% вискоза	AgADAgADXKkxG1ELmUhI575Vveq9pakdMw4ABGYtoREceYDRhnkEAAEC
58	Джемпер с нагрудным карманом	1399	Светло-серый	21	Выполнен из вязаного тонкого меланжевого полотна, круглый вырез горловины, рукава 3/4, состав: 83% полиэстер, 17% вискоза	AgADAgADjagxG2RKoEifbv_cbrWWySMGMw4ABDPtfoBhgDqzGXkEAAEC
59	Джемпер из пряжи шениль	1799	Светло-красный	21	Выполнен из мягкой пряжи шениль, круглый вырез горловины, длинные рукава, состав: 100% полиэстер	AgADAgADXakxG1ELmUgRNPvLj-UBeW0VnA4ABG7VP9F5YHkgkNIBAAEC
60	Джемпер с цветочным принтом	1199	Светло-серый	21	Выполнен из неоднородного полотна, круглый вырез горловины, спущенные рукава, состав: 58% полиэстер, 42% вискоза	AgADAgADXqkxG1ELmUgN2LAWQiCwYoIFnA4ABM-P0u8_SOqKy9QBAAEC
61	Джемпер в полоску	1599	Белый	21	Выполнен из структурного полотна в полоску, короткие рукава, боковые ленты, круглый вырез горловины, состав: 78% полиэстер, 22% вискоза	AgADAgADX6kxG1ELmUhMwr12BVIuqOYPMw4ABPZKuUZjjO6s9nQEAAEC
62	Джемпер с объёмным принтом	1599	Светло-красный	21	Выполнен из гладкого плотного полотна, круглый вырез горловины, длинные рукава, состав: 73% полиэстер, 22% вискоза, 5% эластан	AgADAgADYKkxG1ELmUj4m06C2aQrzeTxmw4ABLd4fbwgr38_U9YBAAEC
63	Джемпер с золотой накаткой	1199	Хаки	21	Выполнен из структурного поливискозного полотна с накаткой, круглый вырез горловины, рукава 3/4, состав: 63% вискоза, 37% полиэстер	AgADAgADYakxG1ELmUhmXnuWPCRnXtGbQw4ABDEpu7G_RteVvNQBAAEC
64	Джемпер с люрексом	1799	Светло-бежевый	21	Выполнен из пряжи с люрексом, круглый вырез горловины, короткие рукава, состав: 46% акрил, 42% полиэстер, 12% металлизированное волокно	AgADAgADYqkxG1ELmUiPnRQg995Ao_Kgmg4ABFlV3uwilNPf99YBAAEC
65	Толстовка с паетками	1599	Хаки	22	Выполнена из структурного полотна, круглый вырез горловины, длинные рукава, состав: 96% полиэстер, 4% эластан	AgADAgADZKkxG1ELmUjxs2_Xh-K0E4f0mw4ABF6syEkso0EMI9UBAAEC
66	Толстовка с капюшоном	1799	Черный	22	Выполнена из неопрена, круглый вырез горловины, длинные рукава, состав: 69% полиэстер, 31% вискоза	AgADAgADZakxG1ELmUiiuo6lzA5qQHYbnA4ABAETFQmAJtXISdMBAAEC
67	Толстовка с принтом "цветы"	1799	Светло-серый	22	Выполнен из меланжевого футерированного полотна с лёгким начёсом, круглый вырез горловины, состав: 60% хлопок, 40% полиэстер	AgADAgADZqkxG1ELmUjgijU1IFhxpVkJMw4ABA3ZybpCORk_inMEAAEC
68	Толстовка с цветочным принтом	1399	Молочный	22	Выполнена из мягкого трикотажа, круглый вырез горловины, длинные рукава, состав: 99% хлопок, 1% вискоза	AgADAgADZ6kxG1ELmUj94uIyOaDZnqkNnA4ABIvXULvWfygcOs8BAAEC
69	Толстовка с вырезами	1399	Светло-розовый	22	Выполнена из мягкой ткани, круглый вырез горловины, открытые плечи, рюши, состав: 60% хлопок, 40% полиэстер	AgADAgADaKkxG1ELmUiV0CdvSpaaUwm2mg4ABAILeICuyXqL7dQBAAEC
70	Толстовка в цветочный принт	1199	Тёмно-синий	22	Выполнена из принтованного футерированного полотна, длинные рукава, принт "цветы", состав: 60% хлопок, 40% полиэстер	AgADAgADaakxG1ELmUgx93bxidsfqvQHnA4ABJJqje11TDcqm9QBAAEC
71	Блузка с открытыми плечами	1199	Голубой	23	Выполнена из хлопкового полотна, открытые плечи, рукава 3/4, состав: 100% хлопок	AgADAgADaqkxG1ELmUiFGB0FgpPoAAFuBZwOAATf8gIJQQWu1bPRAQABAg
72	Хлопковая блузка с вышивкой	1599	Белый	23	Выполнена из вышитого хлопкового полотна, заострённый воротник, длинные рукава, состав:100% хлопок	AgADAgADa6kxG1ELmUg86nF7TSqIiUujmg4ABNIsTQPwNGy0xNkBAAEC
73	Блузка с оборками на рукавах	1199	Молочный	23	Выполнена из мягкого полиэстерового полотна, круглый вырез горловины, спущенные плечи, состав: 100% полиэстер	AgADAgADbKkxG1ELmUhPekVNZaxadPv-Mg4ABDOOivRFS_cm-3MEAAEC
74	Блузка с открытыми плечами	1199	Белый	23	Выполнена из хлопкового полотна, открытые плечи, эластичная резинка по плечам, состав: 100% хлопок	AgADAgADbakxG1ELmUji-3sLGN8rkvwXMw4ABMuzsE3_R1oTpnYEAAEC
75	Блузка с оборкой	1399	Светло-розовый	23	Выполнена из лёгкого хлопкового структурного полотна, круглый вырез горловины, вышивки спереди, рукава 7/8, состав: 60% хлопок, 40% полиэстер	AgADAgADbqkxG1ELmUiKUe9_5c6LOokNnA4ABGDAc1wxcpgtq9QBAAEC
76	Рубашка в ботанический принт	1599	Белый	24	Выполнена из хлопковой ткани, классическая форма воротника, застёжка на пуговицы, состав: 100% вискоза	AgADAgADcKkxG1ELmUhpmXw6DhaqPJsbMw4ABMKPULHr-u0WT3UEAAEC
77	Рубашка из денима	1599	Голубой	24	Выполнена из лёгкого денима, классическая форма воротника, закругленный край, состав: 100% хлопок	AgADAgADcakxG1ELmUiP_zuvApo9meAFnA4ABMo8LmJKfTbSZNIBAAEC
78	Рубашка их хлопка	1999	Белый	24	Выполнена из мерсеризованного хлопка, застёжка-планка, рукава 3/4, состав: 71% хлопок, 29% нейлон	AgADAgADcqkxG1ELmUiD_KYll4dceRQSMw4ABG7bHiLi6LNfbXgEAAEC
79	Объёмная рубашка	1399	Красный	24	Выполнена из хлопковой ткани, застёжка на пуговицы, рукава 3/4, состав: 60% хлопок, 40% полиэстер	AgADAgADc6kxG1ELmUjOYOz-hdRaV_rvAw4ABL7WrVPGcetSqrcDAAEC
80	Рубашка из ткани шамбре	1599	Индиго	24	Выполнена из хлопковой ткани, застёжка на пуговицы, рукава 3/4, состав: 100% хлопок	AgADAgADdKkxG1ELmUhGVSdFT-6SdiMOMw4ABKWQYPnYvfsCMHUEAAEC
81	Рубашка из хлопка	1599	Тёмно-синий	24	Выполнена из хлоковой ткани с эластаном, классическая форма воротника, рукава 3/4, состав: 69% хлопок, 31% нейлон	AgADAgADdakxG1ELmUj4-hLdqh9Vk4S3mg4ABMA0eZjE5Jk07tQBAAEC
82	Рубашка с нашивками	1599	Хаки	24	Выполнена из мягкого тенсельного полотна, застёжка на пуговицы, с эффектом стирки, состав; 100% лиоцелл	AgADAgADdqkxG1ELmUjXXjehiP3oC_ykmg4ABJvAYIBbaqoBpdEBAAEC
83	Рубашка их хлопка	1599	Белый	24	Выполнена из хлопковой ткани, узкая планка с пуговицами спереди, рукава с манжетами 3/4, состав: 69% хлопок, 31% нейлон	AgADAgADd6kxG1ELmUjS_MF7rPP5XXmlmg4ABFWgBrVUbzl-LdcBAAEC
84	Брюки "5 карманов"	1599	Винный	25	Выполнены из эластичного хлопкового твила, узкие брюки, нормальное положение линии талии, состав: 98% хлопок, 2% эластан	AgADAgADeKkxG1ELmUjamqDH-Xcybreamg4ABHyAL3XTZHYk8dMBAAEC
85	Брюки бэгги из лиоцелла	2499	Тёмно-синий	25	Выполнены из шелковистого лиоцелла, зауженный силуэт, эластичный пояс, состав: 100% лиоцелл	AgADAgADj6gxG2RKoEjyhUWWHMJn2zPzmw4ABNSbp2C1gXU5PdEBAAEC
86	Джоггеры	1799	Чёрный	25	Выполнены из плотной поливискозной ткани с эластаном, зауженный к низу силуэт, два боковых кармана, состав: 70% вискоза, 30% нейлон	AgADAgADeakxG1ELmUjOZxle6lv1oIEHnA4ABDkeLEG_VX8b4NQBAAEC
87	Брюки из структурного хлопка	2499	Тёмно-синий	25	Выполнены из эластичного хлопка, брюки со стрелками, зауженный силуэт, состав: 98% хлопок, 2% эластан	AgADAgADeqkxG1ELmUgjb9PVAegfIxUcnA4ABPD-BX9lv7EKntMBAAEC
111	Майка спортивная	2399	Зелёный	28	Состав: 100% полиэстер	AgADAgADKqkxG7uKmEiMo4BoAlfMTzwBMw4ABGUtwfvKH5x4_n4EAAEC
88	Жаккардовые брюки	1999	Чёрный	25	Выполнены из структурной поливискозной ткани, брюки со стрелками, узкие прямые брюки, состав: 70% полиэстер, 30% вискоза	AgADAgADe6kxG1ELmUjPWBqI5dD2ScO3mg4ABJLQq1ktrQsqvdgBAAEC
89	Брюки из структурного хлопка	2499	Красный	25	Выполнены из эластичного хлопка, брюки со стрелками, состав: 98% хлопок, 2% эластан	AgADAgADfKkxG1ELmUjmz4nljMYsiDGcQw4ABMJ9Y00twdRTz9cBAAEC
90	Узкие брюки	1599	Тёмный хаки	25	Выполнены из мягкого хлопкового полотна с тенселем, очень узкий силуэт, линия талии на среднем уровне, состав: 70% хлопок, 30% модал	AgADAgADfakxG1ELmUhmrhcsycVDtDITnA4ABHaJXBEKbCc2RNcBAAEC
91	Текстильные брюки	1799	Чёрный	25	Состав: 70% полиэстер, 28% вискоза, 2% эластан	AgADAgADfqkxG1ELmUg85o_cbpIMhJi4mg4ABDCBhhZrUsYfktIBAAEC
92	Брюки бэгги из крепа	1999	Чёрный	25	Выполнены из плотного крепа, брюки свободные, зауженный к низу силуэт, состав: 100% полиэстер	AgADAgADf6kxG1ELmUjbLihI6cP_nKmhmg4ABGdEFKisKlwtstEBAAEC
93	Брюки-дудочки	1799	Чёрный	25	Выполнены из очень эластичной ткани, узкие брюки, два боковых кармана, состав: 90% полиамид, 10% эластан	AgADAgADJKkxG7uKmEiXCmuaWythQ6YZnA4ABK7oJRygWs0PNNIBAAEC
94	Трикотажные брюки	999	Светло-серый	25	Выполнены из мягкого трикотажа, линия талии на среднем уровне, два боковых кармана, состав: 80% хлопок, 20% полиэстер	AgADAgADgKkxG1ELmUit4zXHYmNYaWP0mw4ABHId80p8BC_vq9IBAAEC
95	Свободные джинсы с вышивкой	2499	Светло-синий	26	Выполнены из денимного полотна, высокое положение линии талии, цветочная вышивка спереди, состав: 100% хлопок	AgADAgADJakxG7uKmEg8czSZChl5pM0VMw4ABKQutG6U5nSX3HUEAAEC
96	Узкие джинсы с потёртостями	1999	Светло-синий	26	Выполнены из эластичной джинсовой ткани, классический дизайн, линия талии на среднем уровне, состав: 99% хлопок, 1% эластан	AgADAgADgakxG1ELmUiS2woPvzW8yQILMw4ABLr6_uJ_J1LvNHUEAAEC
97	Свободные джинсы с принтом	2499	Светло-синий	26	Выполнены из плотного хлопка, свободный силуэт, зауженные к низу штанины, завышенная талия, состав: 100% хлопок	AgADAgADgqkxG1ELmUhWMubwY9a0llOfmg4ABDr-gooeiX7Wd9gBAAEC
98	Джинсы бойфренд с вышивкой	2499	Светло-синий	26	Выполнены из денимной ткани, силуэт с 5 карманами, линия талии на среднем уровне, состав: 100% хлопок	AgADAgADg6kxG1ELmUgy_FfnLcQRCKIRMw4ABA_s4jemr50SI3QEAAEC
99	Футболка-боди	699	Красный	27	Выполнена из эластичного хлопкового джерси, с открытыми плечами, короткие рукава, узкий силуэт, состав: 95% хлопок, 5% эластан	AgADAgADhKkxG1ELmUgEif2Tm_pGMrAGMw4ABNZC31Prdi2FIHoEAAEC
100	Футболка в сетку	1199	Чёрный	27	Выполнена из сетки с внутренней майкой из тонкого джерси, круглый вырез горловины, короткие рукава, состав: 100% полиэстер	AgADAgADKKkxG7uKmEiLUK3Mm89Qhrbzmw4ABPSlqlfhCUTJ7c8BAAEC
101	Футболка с принтом цветы	899	Белый	27	Выполнена из неоднородного мягкого трикотажного полотна, круглый вырез горловины, принт цветы, линия низа изогнутая, состав: 100% полиэстер	AgADAgADhakxG1ELmUhgYE_LvOGHScIFnA4ABN4_omRO91E6BdMBAAEC
102	Футболка с нашивками	599	Жемчужно-белый	27	Выполнена из тонкого джерси, круглый вырез горловины, короткие рукава, яркие нашивки, состав: 100% хлопок	AgADAgADhqkxG1ELmUiDrlDQ4cr5n2ICnA4ABCwieOq8HIL3SdIBAAEC
103	Футболка с принтом	899	Белый	27	Выполнена из тонкого джерси, круглый вырез горловины, короткие рукава с подворотами, открытая спинка, состав: 100% хлопок	AgADAgADh6kxG1ELmUglMmHb2grgFV0TMw4ABMPvBRhZimKWBHYEAAEC
104	Футболка с узлом	799	Хаки	27	Выполнена из неоднородного полотна, круглый вырез горловины, спущенные плечи, дизайн с узлом, слегка свободный силуэт, состав: 100% полиэстер	AgADAgADiKkxG1ELmUiANocO3TUlcNK1mg4ABBWaCkK3Ur6sp88BAAEC
105	Замшевая футболка	999	Розовый	27	Выполнена из мягкого замшевого полотна, круглый вырез горловины, коротие рукава, состав: 100% полиэстер	AgADAgADiakxG1ELmUjxoPBqPzjm36zqAw4ABDYSgw1la4AtrrMDAAEC
106	Камуфляжная футболка	799	Тёмный хаки	27	Выполнена из тонкого джерси, круглый вырез горловины, короткие рукава, камуфляжный принт, спущенные плечи, состав: 100% хлопок	AgADAgADiqkxG1ELmUg1xAcVrkbtbEPvAw4ABKcPR0EUD0uo6K8DAAEC
107	Базовый топ	399	Белый	28	Выполнен из мягкого хлопкового полотна, тонкие бретели, базовый полуприлегающий силуэт, состав:95% хлопок, 5% эластан	AgADAgADi6kxG1ELmUiQcenrj2DgU6sRnA4ABDcKAngi0kEBFdQBAAEC
108	Майка из хлопка	399	Чёрный	28	Выполнена из мягкого хлопкового полотна, широкие бретели, овальный вырез горловины, состав: 95% хлопок, 5% эластан	AgADAgADjKkxG1ELmUiFgNuAyLc_1A2jmg4ABC2rzTaxkhI039IBAAEC
109	Топ из искусственной кожи	1799	Чёрный	28	Выполнен из крепированного полотна, круглый вырез горловины, без рукавов, застёжка на молнии сзади, состав: 100% полиэстер	AgADAgADKakxG7uKmEitVeb81eEvF4QXnA4ABKkyedwNFJIRns4BAAEC
112	Топ спортивный	2799	Красный	28	Состав: 74% полиэстер, 26% эластан	AgADAgADjqkxG1ELmUhsf35XQYDVrxq5mg4ABAvfibS50PWNctIBAAEC
\.


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: tables; Owner: postgres
--

SELECT pg_catalog.setval('products_id_seq', 112, true);


--
-- Data for Name: reviews; Type: TABLE DATA; Schema: tables; Owner: postgres
--

COPY reviews (id, id_product, id_user, date, description) FROM stdin;
\.


--
-- Data for Name: sizes; Type: TABLE DATA; Schema: tables; Owner: postgres
--

COPY sizes (id, title) FROM stdin;
1	XS
2	S
3	M
4	L
5	XL
6	XXL
7	XXXL
8	35
9	36
10	37
11	38
12	39
13	40
14	41
15	42
16	43
17	44
18	45
19	46
\.


--
-- Name: sizes_id_seq; Type: SEQUENCE SET; Schema: tables; Owner: postgres
--

SELECT pg_catalog.setval('sizes_id_seq', 19, true);


--
-- Data for Name: users; Type: TABLE DATA; Schema: tables; Owner: postgres
--

COPY users (id, username, phone, fullname, address, registration_completed, id_current, current_offset) FROM stdin;
294176487	none	79533174973	none	Калуга суворова 13	t	15	0
364794408	none	89997352725	none	Калуга, Суворова 9	t	15	0
197115347	none	none	none	none	f	0	0
326419	none	none	none	none	f	15	0
335028012	none	none	none	none	f	16	0
\.


--
-- Name: catalog catalog_pkey; Type: CONSTRAINT; Schema: tables; Owner: postgres
--

ALTER TABLE ONLY catalog
    ADD CONSTRAINT catalog_pkey PRIMARY KEY (id);


--
-- Name: order_product order_product_pkey; Type: CONSTRAINT; Schema: tables; Owner: postgres
--

ALTER TABLE ONLY order_product
    ADD CONSTRAINT order_product_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: tables; Owner: postgres
--

ALTER TABLE ONLY orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: tables; Owner: postgres
--

ALTER TABLE ONLY products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: tables; Owner: postgres
--

ALTER TABLE ONLY reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- Name: sizes sizes_pkey; Type: CONSTRAINT; Schema: tables; Owner: postgres
--

ALTER TABLE ONLY sizes
    ADD CONSTRAINT sizes_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: tables; Owner: postgres
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: order_product check_items_quantity; Type: TRIGGER; Schema: tables; Owner: postgres
--

CREATE TRIGGER check_items_quantity BEFORE INSERT ON order_product FOR EACH ROW EXECUTE PROCEDURE check_items_quantity();


--
-- PostgreSQL database dump complete
--

