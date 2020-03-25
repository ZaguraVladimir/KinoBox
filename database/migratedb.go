package database

import (
	"database/sql"
	"log"
)

var initQuery = `
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Таблица: countries
DROP TABLE IF EXISTS countries; CREATE TABLE countries (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, name_rus TEXT NOT NULL, wiki TEXT NOT NULL, flag TEXT NOT NULL, iso TEXT (2) NOT NULL, UNIQUE (iso) ON CONFLICT IGNORE);
INSERT INTO countries (id, name, name_rus, wiki, flag, iso) VALUES
(1, 'Australia', 'Австралия', 'https://ru.wikipedia.org/wiki/%D0%90%D0%B2%D1%81%D1%82%D1%80%D0%B0%D0%BB%D0%B8%D1%8F', './img/flags/AU.png', 'AU'),
(2, 'Austria', 'Австрия', 'https://ru.wikipedia.org/wiki/%D0%90%D0%B2%D1%81%D1%82%D1%80%D0%B8%D1%8F', './img/flags/AT.png', 'AT'),
(3, 'Azerbaijan', 'Азербайджан', 'https://ru.wikipedia.org/wiki/%D0%90%D0%B7%D0%B5%D1%80%D0%B1%D0%B0%D0%B9%D0%B4%D0%B6%D0%B0%D0%BD', './img/flags/AZ.png', 'AZ'),
(4, 'Albania', 'Албания', 'https://ru.wikipedia.org/wiki/%D0%90%D0%BB%D0%B1%D0%B0%D0%BD%D0%B8%D1%8F', './img/flags/AL.png', 'AL'),
(5, 'Algeria', 'Алжир', 'https://ru.wikipedia.org/wiki/%D0%90%D0%BB%D0%B6%D0%B8%D1%80', './img/flags/DZ.png', 'DZ'),
(6, 'Argentina', 'Аргентина', 'https://ru.wikipedia.org/wiki/%D0%90%D1%80%D0%B3%D0%B5%D0%BD%D1%82%D0%B8%D0%BD%D0%B0', './img/flags/AR.png', 'AR'),
(7, 'Armenia', 'Армения', 'https://ru.wikipedia.org/wiki/%D0%90%D1%80%D0%BC%D0%B5%D0%BD%D0%B8%D1%8F', './img/flags/AM.png', 'AM'),
(8, 'Afghanistan', 'Афганистан', 'https://ru.wikipedia.org/wiki/%D0%90%D1%84%D0%B3%D0%B0%D0%BD%D0%B8%D1%81%D1%82%D0%B0%D0%BD', './img/flags/AF.png', 'AF'),
(9, 'Bahamas', 'Багамы', 'https://ru.wikipedia.org/wiki/%D0%91%D0%B0%D0%B3%D0%B0%D0%BC%D1%81%D0%BA%D0%B8%D0%B5_%D0%9E%D1%81%D1%82%D1%80%D0%BE%D0%B2%D0%B0', './img/flags/BS.png', 'BS'),
(10, 'Bangladesh', 'Бангладеш', 'https://ru.wikipedia.org/wiki/%D0%91%D0%B0%D0%BD%D0%B3%D0%BB%D0%B0%D0%B4%D0%B5%D1%88', './img/flags/BD.png', 'BD'),
(11, 'Belarus', 'Белоруссия ', 'https://ru.wikipedia.org/wiki/%D0%91%D0%B5%D0%BB%D0%BE%D1%80%D1%83%D1%81%D1%81%D0%B8%D1%8F', './img/flags/BY.png', 'BY'),
(12, 'Belgium', 'Бельгия', 'https://ru.wikipedia.org/wiki/%D0%91%D0%B5%D0%BB%D1%8C%D0%B3%D0%B8%D1%8F', './img/flags/BE.png', 'BE'),
(13, 'Bulgaria', 'Болгария', 'https://ru.wikipedia.org/wiki/%D0%91%D0%BE%D0%BB%D0%B3%D0%B0%D1%80%D0%B8%D1%8F', './img/flags/BG.png', 'BG'),
(14, 'Bolivia', 'Боливия', 'https://ru.wikipedia.org/wiki/%D0%91%D0%BE%D0%BB%D0%B8%D0%B2%D0%B8%D1%8F', './img/flags/BO.png', 'BO'),
(15, 'Bosnia & Herzegovina', 'Босния и Герцеговина', 'https://ru.wikipedia.org/wiki/%D0%91%D0%BE%D1%81%D0%BD%D0%B8%D1%8F_%D0%B8_%D0%93%D0%B5%D1%80%D1%86%D0%B5%D0%B3%D0%BE%D0%B2%D0%B8%D0%BD%D0%B0', './img/flags/BA.png', 'BA'),
(16, 'Botswana', 'Ботсвана', 'https://ru.wikipedia.org/wiki/%D0%91%D0%BE%D1%82%D1%81%D0%B2%D0%B0%D0%BD%D0%B0', './img/flags/BW.png', 'BW'),
(17, 'Brazil', 'Бразилия', 'https://ru.wikipedia.org/wiki/%D0%91%D1%80%D0%B0%D0%B7%D0%B8%D0%BB%D0%B8%D1%8F', './img/flags/BR.png', 'BR'),
(18, 'Bhutan', 'Бутан', 'https://ru.wikipedia.org/wiki/%D0%91%D1%83%D1%82%D0%B0%D0%BD', './img/flags/BT.png', 'BT'),
(19, 'United Kingdom', 'Великобритания', 'https://ru.wikipedia.org/wiki/%D0%92%D0%B5%D0%BB%D0%B8%D0%BA%D0%BE%D0%B1%D1%80%D0%B8%D1%82%D0%B0%D0%BD%D0%B8%D1%8F', './img/flags/GB.png', 'GB'),
(20, 'Hungary', 'Венгрия', 'https://ru.wikipedia.org/wiki/%D0%92%D0%B5%D0%BD%D0%B3%D1%80%D0%B8%D1%8F', './img/flags/HU.png', 'HU'),
(21, 'Venezuela', 'Венесуэла', 'https://ru.wikipedia.org/wiki/%D0%92%D0%B5%D0%BD%D0%B5%D1%81%D1%83%D1%8D%D0%BB%D0%B0', './img/flags/VE.png', 'VE'),
(22, 'Viet Nam', 'Вьетнам', 'https://ru.wikipedia.org/wiki/%D0%92%D1%8C%D0%B5%D1%82%D0%BD%D0%B0%D0%BC', './img/flags/VN.png', 'VN'),
(23, 'Ghana', 'Гана', 'https://ru.wikipedia.org/wiki/%D0%93%D0%B0%D0%BD%D0%B0', './img/flags/GH.png', 'GH'),
(24, 'Guinea', 'Гвинея', 'https://ru.wikipedia.org/wiki/%D0%93%D0%B2%D0%B8%D0%BD%D0%B5%D1%8F', './img/flags/GN.png', 'GN'),
(25, 'Germany', 'Германия', 'https://ru.wikipedia.org/wiki/%D0%93%D0%B5%D1%80%D0%BC%D0%B0%D0%BD%D0%B8%D1%8F', './img/flags/DE.png', 'DE'),
(26, 'Hong Kong', 'Гонконг', 'https://ru.wikipedia.org/wiki/%D0%93%D0%BE%D0%BD%D0%BA%D0%BE%D0%BD%D0%B3', './img/flags/HK.png', 'HK'),
(27, 'Greece', 'Греция', 'https://ru.wikipedia.org/wiki/%D0%93%D1%80%D0%B5%D1%86%D0%B8%D1%8F', './img/flags/GR.png', 'GR'),
(28, 'Georgia', 'Грузия', 'https://ru.wikipedia.org/wiki/%D0%93%D1%80%D1%83%D0%B7%D0%B8%D1%8F', './img/flags/GE.png', 'GE'),
(29, 'Denmark', 'Дания', 'https://ru.wikipedia.org/wiki/%D0%94%D0%B0%D0%BD%D0%B8%D1%8F', './img/flags/DK.png', 'DK'),
(30, 'Egypt', 'Египет', 'https://ru.wikipedia.org/wiki/%D0%95%D0%B3%D0%B8%D0%BF%D0%B5%D1%82', './img/flags/EG.png', 'EG'),
(31, 'Israel', 'Израиль', 'https://ru.wikipedia.org/wiki/%D0%98%D0%B7%D1%80%D0%B0%D0%B8%D0%BB%D1%8C', './img/flags/IL.png', 'IL'),
(32, 'India', 'Индия', 'https://ru.wikipedia.org/wiki/%D0%98%D0%BD%D0%B4%D0%B8%D1%8F', './img/flags/IN.png', 'IN'),
(33, 'Indonesia', 'Индонезия', 'https://ru.wikipedia.org/wiki/%D0%98%D0%BD%D0%B4%D0%BE%D0%BD%D0%B5%D0%B7%D0%B8%D1%8F', './img/flags/ID.png', 'ID'),
(34, 'Jordan', 'Иордания', 'https://ru.wikipedia.org/wiki/%D0%98%D0%BE%D1%80%D0%B4%D0%B0%D0%BD%D0%B8%D1%8F', './img/flags/JO.png', 'JO'),
(35, 'Iraq', 'Ирак', 'https://ru.wikipedia.org/wiki/%D0%98%D1%80%D0%B0%D0%BA', './img/flags/IQ.png', 'IQ'),
(36, 'Iran', 'Иран', 'https://ru.wikipedia.org/wiki/%D0%98%D1%80%D0%B0%D0%BD', './img/flags/IR.png', 'IR'),
(37, 'Ireland', 'Ирландия', 'https://ru.wikipedia.org/wiki/%D0%98%D1%80%D0%BB%D0%B0%D0%BD%D0%B4%D0%B8%D1%8F', './img/flags/IE.png', 'IE'),
(38, 'Iceland', 'Исландия', 'https://ru.wikipedia.org/wiki/%D0%98%D1%81%D0%BB%D0%B0%D0%BD%D0%B4%D0%B8%D1%8F', './img/flags/IS.png', 'IS'),
(39, 'Spain', 'Испания', 'https://ru.wikipedia.org/wiki/%D0%98%D1%81%D0%BF%D0%B0%D0%BD%D0%B8%D1%8F', './img/flags/ES.png', 'ES'),
(40, 'Italy', 'Италия', 'https://ru.wikipedia.org/wiki/%D0%98%D1%82%D0%B0%D0%BB%D0%B8%D1%8F', './img/flags/IT.png', 'IT'),
(41, 'Kazakhstan', 'Казахстан', 'https://ru.wikipedia.org/wiki/%D0%9A%D0%B0%D0%B7%D0%B0%D1%85%D1%81%D1%82%D0%B0%D0%BD', './img/flags/KZ.png', 'KZ'),
(42, 'Cambodia', 'Камбоджа', 'https://ru.wikipedia.org/wiki/%D0%9A%D0%B0%D0%BC%D0%B1%D0%BE%D0%B4%D0%B6%D0%B0', './img/flags/KH.png', 'KH'),
(43, 'Cameroon', 'Камерун', 'https://ru.wikipedia.org/wiki/%D0%9A%D0%B0%D0%BC%D0%B5%D1%80%D1%83%D0%BD', './img/flags/CM.png', 'CM'),
(44, 'Canada', 'Канада', 'https://ru.wikipedia.org/wiki/%D0%9A%D0%B0%D0%BD%D0%B0%D0%B4%D0%B0', './img/flags/CA.png', 'CA'),
(45, 'Kyrgyzstan', 'Киргизия', 'https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B3%D0%B8%D0%B7%D0%B8%D1%8F', './img/flags/KG.png', 'KG'),
(46, 'China', 'Китай', 'https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%82%D0%B0%D0%B9%D1%81%D0%BA%D0%B0%D1%8F_%D0%9D%D0%B0%D1%80%D0%BE%D0%B4%D0%BD%D0%B0%D1%8F_%D0%A0%D0%B5%D1%81%D0%BF%D1%83%D0%B1%D0%BB%D0%B8%D0%BA%D0%B0', './img/flags/CN.png', 'CN'),
(47, 'Colombia', 'Колумбия', 'https://ru.wikipedia.org/wiki/%D0%9A%D0%BE%D0%BB%D1%83%D0%BC%D0%B1%D0%B8%D1%8F', './img/flags/CO.png', 'CO'),
(48, 'Korea, D.P.R.', 'Северная Корея', 'https://ru.wikipedia.org/wiki/%D0%9A%D0%BE%D1%80%D0%B5%D0%B9%D1%81%D0%BA%D0%B0%D1%8F_%D0%9D%D0%B0%D1%80%D0%BE%D0%B4%D0%BD%D0%BE-%D0%94%D0%B5%D0%BC%D0%BE%D0%BA%D1%80%D0%B0%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B0%D1%8F_%D0%A0%D0%B5%D1%81%D0%BF%D1%83%D0%B1%D0%BB%D0%B8%D0%BA%D0%B0', './img/flags/ KP.png', ' KP'),
(49, 'Korea', 'Южная Корея', 'https://ru.wikipedia.org/wiki/%D0%A0%D0%B5%D1%81%D0%BF%D1%83%D0%B1%D0%BB%D0%B8%D0%BA%D0%B0_%D0%9A%D0%BE%D1%80%D0%B5%D1%8F', './img/flags/KR.png', 'KR'),
(50, 'Costa Rica', 'Коста-Рика', 'https://ru.wikipedia.org/wiki/%D0%9A%D0%BE%D1%81%D1%82%D0%B0-%D0%A0%D0%B8%D0%BA%D0%B0', './img/flags/CR.png', 'CR'),
(51, 'Cuba', 'Куба', 'https://ru.wikipedia.org/wiki/%D0%9A%D1%83%D0%B1%D0%B0', './img/flags/CU.png', 'CU'),
(52, 'Kuwait', 'Кувейт', 'https://ru.wikipedia.org/wiki/%D0%9A%D1%83%D0%B2%D0%B5%D0%B9%D1%82', './img/flags/KW.png', 'KW'),
(53, 'Lao P.D.R.', 'Лаос', 'https://ru.wikipedia.org/wiki/%D0%9B%D0%B0%D0%BE%D1%81', './img/flags/LA.png', 'LA'),
(54, 'Latvia', 'Латвия', 'https://ru.wikipedia.org/wiki/%D0%9B%D0%B0%D1%82%D0%B2%D0%B8%D1%8F', './img/flags/LV.png', 'LV'),
(55, 'Lebanon', 'Ливан', 'https://ru.wikipedia.org/wiki/%D0%9B%D0%B8%D0%B2%D0%B0%D0%BD', './img/flags/LB.png', 'LB'),
(56, 'Libyan Arab Jamahiriya', 'Ливия', 'https://ru.wikipedia.org/wiki/%D0%9B%D0%B8%D0%B2%D0%B8%D1%8F', './img/flags/LY.png', 'LY'),
(57, 'Lithuania', 'Литва', 'https://ru.wikipedia.org/wiki/%D0%9B%D0%B8%D1%82%D0%B2%D0%B0', './img/flags/LT.png', 'LT'),
(58, 'Liechtenstein', 'Лихтенштейн', 'https://ru.wikipedia.org/wiki/%D0%9B%D0%B8%D1%85%D1%82%D0%B5%D0%BD%D1%88%D1%82%D0%B5%D0%B9%D0%BD', './img/flags/LI.png', 'LI'),
(59, 'Luxembourg', 'Люксембург', 'https://ru.wikipedia.org/wiki/%D0%9B%D1%8E%D0%BA%D1%81%D0%B5%D0%BC%D0%B1%D1%83%D1%80%D0%B3', './img/flags/LU.png', 'LU'),
(60, 'Mauritania', 'Мавритания', 'https://ru.wikipedia.org/wiki/%D0%9C%D0%B0%D0%B2%D1%80%D0%B8%D1%82%D0%B0%D0%BD%D0%B8%D1%8F', './img/flags/MR.png', 'MR'),
(61, 'Macedonia', 'Македония', 'https://ru.wikipedia.org/wiki/%D0%A0%D0%B5%D1%81%D0%BF%D1%83%D0%B1%D0%BB%D0%B8%D0%BA%D0%B0_%D0%9C%D0%B0%D0%BA%D0%B5%D0%B4%D0%BE%D0%BD%D0%B8%D1%8F', './img/flags/MK.png', 'MK'),
(62, 'Malta', 'Мальта', 'https://ru.wikipedia.org/wiki/%D0%9C%D0%B0%D0%BB%D1%8C%D1%82%D0%B0', './img/flags/MT.png', 'MT'),
(63, 'Morocco', 'Марокко', 'https://ru.wikipedia.org/wiki/%D0%9C%D0%B0%D1%80%D0%BE%D0%BA%D0%BA%D0%BE', './img/flags/MA.png', 'MA'),
(64, 'Mexico', 'Мексика', 'https://ru.wikipedia.org/wiki/%D0%9C%D0%B5%D0%BA%D1%81%D0%B8%D0%BA%D0%B0', './img/flags/MX.png', 'MX'),
(65, 'Moldova', 'Молдавия', 'https://ru.wikipedia.org/wiki/%D0%9C%D0%BE%D0%BB%D0%B4%D0%B0%D0%B2%D0%B8%D1%8F', './img/flags/MD.png', 'MD'),
(66, 'Monaco', 'Монако', 'https://ru.wikipedia.org/wiki/%D0%9C%D0%BE%D0%BD%D0%B0%D0%BA%D0%BE', './img/flags/MC.png', 'MC'),
(67, 'Mongolia', 'Монголия', 'https://ru.wikipedia.org/wiki/%D0%9C%D0%BE%D0%BD%D0%B3%D0%BE%D0%BB%D0%B8%D1%8F', './img/flags/MN.png', 'MN'),
(68, 'Namibia', 'Намибия', 'https://ru.wikipedia.org/wiki/%D0%9D%D0%B0%D0%BC%D0%B8%D0%B1%D0%B8%D1%8F', './img/flags/NA.png', 'NA'),
(69, 'Nepal', 'Непал', 'https://ru.wikipedia.org/wiki/%D0%9D%D0%B5%D0%BF%D0%B0%D0%BBНигер ', './img/flags/NP.png', 'NP'),
(70, 'Nigeria', 'Нигерия ', 'https://ru.wikipedia.org/wiki/%D0%9D%D0%B8%D0%B3%D0%B5%D1%80%D0%B8%D1%8F', './img/flags/NE.png', 'NE'),
(71, 'Netherlands', 'Нидерланды', 'https://ru.wikipedia.org/wiki/%D0%9D%D0%B8%D0%B4%D0%B5%D1%80%D0%BB%D0%B0%D0%BD%D0%B4%D1%8B', './img/flags/NL.png', 'NL'),
(72, 'New Zealand', 'Новая Зеландия', 'https://ru.wikipedia.org/wiki/%D0%9D%D0%BE%D0%B2%D0%B0%D1%8F_%D0%97%D0%B5%D0%BB%D0%B0%D0%BD%D0%B4%D0%B8%D1%8F', './img/flags/NZ.png', 'NZ'),
(73, 'Norway', 'Норвегия', 'https://ru.wikipedia.org/wiki/%D0%9D%D0%BE%D1%80%D0%B2%D0%B5%D0%B3%D0%B8%D1%8F', './img/flags/NO.png', 'NO'),
(74, 'Pakistan', 'Пакистан', 'https://ru.wikipedia.org/wiki/%D0%9F%D0%B0%D0%BA%D0%B8%D1%81%D1%82%D0%B0%D0%BD', './img/flags/PK.png', 'PK'),
(75, 'Panama', 'Панама', 'https://ru.wikipedia.org/wiki/%D0%9F%D0%B0%D0%BD%D0%B0%D0%BC%D0%B0', './img/flags/PA.png', 'PA'),
(76, 'Paraguay', 'Парагвай', 'https://ru.wikipedia.org/wiki/%D0%9F%D0%B0%D1%80%D0%B0%D0%B3%D0%B2%D0%B0%D0%B9', './img/flags/PY.png', 'PY'),
(77, 'Peru', 'Перу', 'https://ru.wikipedia.org/wiki/%D0%9F%D0%B5%D1%80%D1%83', './img/flags/PE.png', 'PE'),
(78, 'Poland', 'Польша', 'https://ru.wikipedia.org/wiki/%D0%9F%D0%BE%D0%BB%D1%8C%D1%88%D0%B0', './img/flags/PL.png', 'PL'),
(79, 'Portugal', 'Португалия', 'https://ru.wikipedia.org/wiki/%D0%9F%D0%BE%D1%80%D1%82%D1%83%D0%B3%D0%B0%D0%BB%D0%B8%D1%8F', './img/flags/PT.png', 'PT'),
(80, 'Russia', 'Россия', 'https://ru.wikipedia.org/wiki/%D0%A0%D0%BE%D1%81%D1%81%D0%B8%D1%8F', './img/flags/RU.png', 'RU'),
(81, 'Romania', 'Румыния', 'https://ru.wikipedia.org/wiki/%D0%A0%D1%83%D0%BC%D1%8B%D0%BD%D0%B8%D1%8F', './img/flags/RO.png', 'RO'),
(82, 'Saudi Arabia', 'Саудовская Аравия', 'https://ru.wikipedia.org/wiki/%D0%A1%D0%B0%D1%83%D0%B4%D0%BE%D0%B2%D1%81%D0%BA%D0%B0%D1%8F_%D0%90%D1%80%D0%B0%D0%B2%D0%B8%D1%8F', './img/flags/SA.png', 'SA'),
(83, 'Senegal', 'Сенегал', 'https://ru.wikipedia.org/wiki/%D0%A1%D0%B5%D0%BD%D0%B5%D0%B3%D0%B0%D0%BB', './img/flags/SN.png', 'SN'),
(84, 'Serbia', 'Сербия', 'https://ru.wikipedia.org/wiki/%D0%A1%D0%B5%D1%80%D0%B1%D0%B8%D1%8F', './img/flags/RS.png', 'RS'),
(85, 'Syrian Arab Republic', 'Сирия', 'https://ru.wikipedia.org/wiki/%D0%A1%D0%B8%D1%80%D0%B8%D1%8F', './img/flags/SY.png', 'SY'),
(86, 'Slovakia', 'Словакия', 'https://ru.wikipedia.org/wiki/%D0%A1%D0%BB%D0%BE%D0%B2%D0%B0%D0%BA%D0%B8%D1%8F', './img/flags/SK.png', 'SK'),
(87, 'Slovenia', 'Словения', 'https://ru.wikipedia.org/wiki/%D0%A1%D0%BB%D0%BE%D0%B2%D0%B5%D0%BD%D0%B8%D1%8F', './img/flags/SI.png', 'SI'),
(88, 'USSR', 'СССР', 'https://ru.wikipedia.org/wiki/%D0%A1%D0%BE%D1%8E%D0%B7_%D0%A1%D0%BE%D0%B2%D0%B5%D1%82%D1%81%D0%BA%D0%B8%D1%85_%D0%A1%D0%BE%D1%86%D0%B8%D0%B0%D0%BB%D0%B8%D1%81%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D1%85_%D0%A0%D0%B5%D1%81%D0%BF%D1%83%D0%B1%D0%BB%D0%B8%D0%BA', './img/flags/SU.png', 'SU'),
(89, 'Sudan', 'Судан', 'https://ru.wikipedia.org/wiki/%D0%A1%D1%83%D0%B4%D0%B0%D0%BD', './img/flags/SD.png', 'SD'),
(90, 'USA', 'США', 'https://ru.wikipedia.org/wiki/%D0%A1%D0%BE%D0%B5%D0%B4%D0%B8%D0%BD%D1%91%D0%BD%D0%BD%D1%8B%D0%B5_%D0%A8%D1%82%D0%B0%D1%82%D1%8B_%D0%90%D0%BC%D0%B5%D1%80%D0%B8%D0%BA%D0%B8', './img/flags/US.png', 'US'),
(91, 'Tajikistan', 'Таджикистан', 'https://ru.wikipedia.org/wiki/%D0%A2%D0%B0%D0%B4%D0%B6%D0%B8%D0%BA%D0%B8%D1%81%D1%82%D0%B0%D0%BD', './img/flags/TJ.png', 'TJ'),
(92, 'Taiwan', 'Тайвань', 'https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%82%D0%B0%D0%B9%D1%81%D0%BA%D0%B0%D1%8F_%D0%A0%D0%B5%D1%81%D0%BF%D1%83%D0%B1%D0%BB%D0%B8%D0%BA%D0%B0', './img/flags/TW.png', 'TW'),
(93, 'Tunisia', 'Тунис', 'https://ru.wikipedia.org/wiki/%D0%A2%D1%83%D0%BD%D0%B8%D1%81', './img/flags/TN.png', 'TN'),
(94, 'Turkmenistan', 'Туркменистан', 'https://ru.wikipedia.org/wiki/%D0%A2%D1%83%D1%80%D0%BA%D0%BC%D0%B5%D0%BD%D0%B8%D1%8F', './img/flags/TM.png', 'TM'),
(95, 'Turkey', 'Турция', 'https://ru.wikipedia.org/wiki/%D0%A2%D1%83%D1%80%D1%86%D0%B8%D1%8F', './img/flags/TR.png', 'TR'),
(96, 'Uzbekistan', 'Узбекистан', 'https://ru.wikipedia.org/wiki/%D0%A3%D0%B7%D0%B1%D0%B5%D0%BA%D0%B8%D1%81%D1%82%D0%B0%D0%BD', './img/flags/UZ.png', 'UZ'),
(97, 'Ukraine', 'Украина', 'https://ru.wikipedia.org/wiki/%D0%A3%D0%BA%D1%80%D0%B0%D0%B8%D0%BD%D0%B0', './img/flags/UA.png', 'UA'),
(98, 'Uruguay', 'Уругвай', 'https://ru.wikipedia.org/wiki/%D0%A3%D1%80%D1%83%D0%B3%D0%B2%D0%B0%D0%B9', './img/flags/UY.png', 'UY'),
(99, 'Philippines', 'Филиппины', 'https://ru.wikipedia.org/wiki/%D0%A4%D0%B8%D0%BB%D0%B8%D0%BF%D0%BF%D0%B8%D0%BD%D1%8B', './img/flags/PH.png', 'PH'),
(100, 'Finland', 'Финляндия', 'https://ru.wikipedia.org/wiki/%D0%A4%D0%B8%D0%BD%D0%BB%D1%8F%D0%BD%D0%B4%D0%B8%D1%8F', './img/flags/FI.png', 'FI'),
(101, 'France', 'Франция', 'https://ru.wikipedia.org/wiki/%D0%A4%D1%80%D0%B0%D0%BD%D1%86%D0%B8%D1%8F', './img/flags/FR.png', 'FR'),
(102, 'Croatia', 'Хорватия', 'https://ru.wikipedia.org/wiki/%D0%A5%D0%BE%D1%80%D0%B2%D0%B0%D1%82%D0%B8%D1%8F', './img/flags/HR.png', 'HR'),
(103, 'Chad', 'Чад', 'https://ru.wikipedia.org/wiki/%D0%A7%D0%B0%D0%B4', './img/flags/TD.png', 'TD'),
(104, 'Montenegro', 'Черногория', 'https://ru.wikipedia.org/wiki/%D0%A7%D0%B5%D1%80%D0%BD%D0%BE%D0%B3%D0%BE%D1%80%D0%B8%D1%8F', './img/flags/ME.png', 'ME'),
(105, 'Czech Republic', 'Чехия', 'https://ru.wikipedia.org/wiki/%D0%A7%D0%B5%D1%85%D0%B8%D1%8F', './img/flags/CZ.png', 'CZ'),
(106, 'Chile', 'Чили', 'https://ru.wikipedia.org/wiki/%D0%A7%D0%B8%D0%BB%D0%B8', './img/flags/CL.png', 'CL'),
(107, 'Switzerland', 'Швейцария', 'https://ru.wikipedia.org/wiki/%D0%A8%D0%B2%D0%B5%D0%B9%D1%86%D0%B0%D1%80%D0%B8%D1%8F', './img/flags/CH.png', 'CH'),
(108, 'Sweden', 'Швеция', 'https://ru.wikipedia.org/wiki/%D0%A8%D0%B2%D0%B5%D1%86%D0%B8%D1%8F', './img/flags/SE.png', 'SE'),
(109, 'Sri Lanka', 'Шри-Ланка', 'https://ru.wikipedia.org/wiki/%D0%A8%D1%80%D0%B8-%D0%9B%D0%B0%D0%BD%D0%BA%D0%B0', './img/flags/LK.png', 'LK'),
(110, 'Estonia', 'Эстония', 'https://ru.wikipedia.org/wiki/%D0%AD%D0%BA%D0%B2%D0%B0%D0%B4%D0%BE%D1%80', './img/flags/EE.png', 'EE'),
(111, 'South Africa', 'ЮАР', 'https://ru.wikipedia.org/wiki/%D0%AE%D0%B6%D0%BD%D0%BE-%D0%90%D1%84%D1%80%D0%B8%D0%BA%D0%B0%D0%BD%D1%81%D0%BA%D0%B0%D1%8F_%D0%A0%D0%B5%D1%81%D0%BF%D1%83%D0%B1%D0%BB%D0%B8%D0%BA%D0%B0', './img/flags/ZA.png', 'ZA'),
(112, 'Jamaica', 'Ямайка', 'https://ru.wikipedia.org/wiki/%D0%AF%D0%BC%D0%B0%D0%B9%D0%BA%D0%B0', './img/flags/JM.png', 'JM'),
(113, 'Japan', 'Япония', 'https://ru.wikipedia.org/wiki/%D0%AF%D0%BF%D0%BE%D0%BD%D0%B8%D1%8F', './img/flags/JP.png', 'JP');

-- Таблица: countries_films
DROP TABLE IF EXISTS countries_films; CREATE TABLE countries_films (film_id INTEGER REFERENCES films (id) NOT NULL, country_id INTEGER REFERENCES countries (id) NOT NULL, UNIQUE (film_id, country_id) ON CONFLICT IGNORE);

-- Таблица: films
DROP TABLE IF EXISTS films; CREATE TABLE films (id INTEGER PRIMARY KEY AUTOINCREMENT, site_id INTEGER REFERENCES sites (id) NOT NULL, name TEXT NOT NULL, name_rus TEXT NOT NULL, year TEXT DEFAULT (0), ref TEXT NOT NULL, quality_id INTEGER NOT NULL REFERENCES quality (id), duration INTEGER DEFAULT (0), serial INTEGER DEFAULT (0), season INTEGER DEFAULT (0), episode INTEGER DEFAULT (0));

-- Таблица: genres
DROP TABLE IF EXISTS genres; CREATE TABLE genres (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE);
INSERT INTO genres (id, name) VALUES (1, 'Биография'),(2, 'Боевик'),(3, 'Вестерн'),(4, 'Военный'),(5, 'Детектив'),(6, 'Детский'),(7, 'Документальный'),(8, 'Драма'),(9, 'Исторический'),(10, 'Комедия'),(11, 'Криминал'),(12, 'Мелодрама'),(13, 'Мистика'),(14, 'Мультфильм'),(15, 'Мюзикл'),(16, 'Отечественный'),(17, 'Приключения'),(18, 'Семейный'),(19, 'Cпортивный'),(20, 'Триллер'),(21, 'Ужасы'),(22, 'Фантастика'),(23, 'Фэнтези');

-- Таблица: genres_films
DROP TABLE IF EXISTS genres_films; CREATE TABLE genres_films (film_id INTEGER REFERENCES films (id) NOT NULL, genre_id INTEGER REFERENCES genres (id) NOT NULL, UNIQUE (film_id, genre_id) ON CONFLICT IGNORE);

-- Таблица: posters_films
DROP TABLE IF EXISTS posters_films; CREATE TABLE posters_films (film_id INTEGER REFERENCES films (id) NOT NULL, ref TEXT NOT NULL, UNIQUE (film_id, ref) ON CONFLICT IGNORE);

-- Таблица: quality
DROP TABLE IF EXISTS quality; CREATE TABLE quality (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, rate INTEGER NOT NULL);
INSERT INTO quality (id, name, rate) VALUES (1, 'Blu-Ray', 1),(2, 'BDRemux', 1),(3, 'HD DVD', 1),(4, 'HD DVDRemux', 1),(5, 'WEB-DL', 1),(6, 'HDTV', 1),(7, 'BDRip', 2),(8, 'HD DVDRip', 2),(9, 'WEB-DLRip', 2),(10, 'HDTVRip', 2),(11, 'HDRip', 2),(12, 'DVD Custom', 3),(13, 'DVD', 3),(14, 'DVDRip', 3),(15, 'PDTV', 4),(16, 'PDTVRip', 4),(17, 'SATRip', 4),(18, 'DVBRip', 4),(19, 'DTVRip', 4),(20, 'TVRip', 4),(21, 'LDRip', 4),(22, 'TC', 5),(23, 'SCR', 5),(24, 'VHSScr', 5),(25, 'DVDScr', 5),(26, 'WP', 5),(27, 'VHSRip', 5),(28, 'TS', 6),(29, 'CAMRip', 6);

-- Таблица: sites
DROP TABLE IF EXISTS sites; CREATE TABLE sites (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, ref TEXT NOT NULL, image_ref TEXT, rate INTEGER NOT NULL, parsing_rules TEXT NOT NULL DEFAULT "", UNIQUE (name, ref) ON CONFLICT IGNORE);

-- Триггер: after_delete_country
CREATE TRIGGER after_delete_country AFTER DELETE ON countries BEGIN DELETE FROM countries_films WHERE country_id = OLD.id; END;

-- Триггер: after_delete_films
CREATE TRIGGER after_delete_films AFTER DELETE ON films BEGIN DELETE FROM genres_films WHERE film_id = OLD.id; DELETE FROM posters_films WHERE film_id = OLD.id; DELETE FROM countries_films WHERE film_id = OLD.id; END;

-- Триггер: after_delete_genre
CREATE TRIGGER after_delete_genre AFTER DELETE ON genres BEGIN DELETE FROM genres_films WHERE genre_id = OLD.id; END;

-- Триггер: after_delete_site
CREATE TRIGGER after_delete_site AFTER DELETE ON sites BEGIN DELETE FROM films WHERE site_id = OLD.id; DELETE FROM images WHERE id = OLD.image_ref; END;

-- Представление: countries_view
CREATE VIEW countries_view AS SELECT name_rus AS country, wiki, flag FROM countries;

COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
`

func migration(db *sql.DB, logger *log.Logger) {
	logger.Println("Выполняется инициализация базы данных")
	if _, err := db.Exec(initQuery); err != nil {
		logger.Fatalf("При инициализации базы данных произошел сбой. Ошибка: %q\r\n", err)
	}
	logger.Println("Инициализация базы данных завершена")
}
