INSERT INTO author (
  id,
  name
) values (1,'Unknown')
ON CONFLICT (id)
DO NOTHING;

--Some 'organ' and 'tcm organ' overlap, so call all 'organ'
INSERT INTO category_type (id, name) VALUES (1, 'organ')
ON CONFLICT (name)
DO NOTHING;

INSERT INTO category_type (id, name) VALUES (2, 'herb peer relation')
ON CONFLICT (name)
DO NOTHING;

INSERT INTO category_type (id, name) VALUES (4, 'language')
ON CONFLICT (name)
DO NOTHING;

INSERT INTO category_type (id, name) VALUES (5, 'pathology description')
ON CONFLICT (name)
DO NOTHING;

INSERT INTO category_type (id, name) VALUES (6, 'treatment verb')
ON CONFLICT (name)
DO NOTHING;

INSERT INTO category_type (id, name) VALUES (7, 'symptom description')
ON CONFLICT (name)
DO NOTHING;

------

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(701, 7, 'dry', '干', 'g')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(702, 7, 'red', '红', 'h')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(703, 7, 'painful', '痛', 't')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(704, 7, 'cough', '咳嗽', 'ks')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(705, 7, 'phelem', '痰', 't')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(601, 6, 'supply', '补', 'b')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(602, 6, 'decrease/lower', '降', 'j')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(603, 6, 'remove', '消、除、清', 'xcq')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(604, 6, 'harmonize', '和', 'h')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(605, 6, 'warm up', '温', 'w')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(501, 5, 'deficient', '虚', 'x')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(502, 5, 'excessive', '实', 'k')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(503, 5, 'cold', '寒', 'h')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(504, 5, 'hot', '热', 'r')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(401, 4, 'English', '', '')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category 
(id, category_type_id, name, name_zh, pinyin) VALUES 
(402, 4, 'Chinese', '', '')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO source
(id, author_id, language_id, name)
values
(1, 1,401, 'Unkown')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (201, 2, 'poisonous', '有毒', 'yd')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (202, 2, 'weaken effects', '抵消药性', 'dxyx')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (101, 1, 'face', '脸', 'l')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (102, 1, 'mouth', '口', 'k')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (103, 1, 'lips', '唇', 'c')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (104, 1, 'hands', '手', 's')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (105, 1, 'feet', '脚', 'j')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (106, 1, 'head', '头', 't')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (107, 1, 'eyes', '眼', 'y')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (108, 1, 'nose', '鼻', 'b')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (109, 1, 'tongue', '舌', 's')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (110, 1, 'throat', '喉', 'h')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (111, 1, 'ears', '耳', 'e')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (112, 1, 'hair (head top)', '头发', 'tf')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (113, 1, 'stomach', '腹', 'f')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (114, 1, 'chest', '胸', 'x')
ON CONFLICT (id)
DO NOTHING;


INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (115, 1, 'yin', '阴', 'y')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (116, 1, 'yang', '阳', 'y')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO category (id, category_type_id, name, name_zh, pinyin) VALUES (117, 1, '', NULL, NULL)
ON CONFLICT (id)
DO NOTHING;