CREATE TABLE IF NOT EXISTS users
(
  id SERIAL PRIMARY KEY,
  name varchar(100) NOT NULL,
  password varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS category_type
(
  id INTEGER PRIMARY KEY,
  name varchar(100) NOT NULL,
  CONSTRAINT unique_cat_type_name UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS category
(
  id INTEGER PRIMARY KEY,
  category_type_id INTEGER REFERENCES category_type(id) NOT NULL,
  name varchar(100) NOT NULL,
  name_zh varchar(100),
  pinyin varchar(50)
);

CREATE TABLE IF NOT EXISTS herb
(
  id SERIAL PRIMARY KEY,
  name varchar(100),
  name_zh varchar(100),
  pinyin varchar(50)
);

CREATE TABLE IF NOT EXISTS herb_peer
(
  herb_id INTEGER REFERENCES herb(id) NOT NULL,
  herb2_id INTEGER REFERENCES herb(id) NOT NULL,
  herb_relation_id INTEGER REFERENCES category(id) NOT NULL,
  PRIMARY KEY (herb_id, herb2_id)
);

CREATE TABLE IF NOT EXISTS author
(
  id SERIAL PRIMARY KEY,
  name varchar(60) NOT NULL,
  CONSTRAINT unique_author UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS source
(
  id SERIAL PRIMARY KEY,
  author_id INTEGER REFERENCES author(id) NOT NULL,
  language_id INTEGER REFERENCES category(id) NOT NULL,
  name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS recipe
(
  id SERIAL PRIMARY KEY,
  name varchar(255),
  name_zh varchar(255),
  pinyin varchar(50),
  source_id INTEGER REFERENCES source(id)
);

CREATE TABLE IF NOT EXISTS recipe_herb
(
  recipe_id INTEGER REFERENCES recipe(id) NOT NULL,
  herb_id INTEGER REFERENCES herb(id) NOT NULL,
  gram INTEGER,
  note varchar(255),
  PRIMARY KEY (recipe_id, herb_id)
);

CREATE TABLE IF NOT EXISTS organ_symptom
(
  id SERIAL PRIMARY KEY,
  organ_id INTEGER REFERENCES category(id),
  symptom_id INTEGER REFERENCES category(id) NOT NULL,
  CONSTRAINT unique_organ_symptom UNIQUE(organ_id, symptom_id)
);

CREATE TABLE IF NOT EXISTS disease
(
  id SERIAL PRIMARY KEY,
  name varchar(150),  
  name_zh varchar(255),
  source_id INTEGER REFERENCES source(id) NOT NULL,
  pinyin varchar(100)
);

CREATE TABLE IF NOT EXISTS disease_organsymptom
(
  disease_id INTEGER REFERENCES disease(id) NOT NULL,
  organ_symptom_id INTEGER REFERENCES organ_symptom(id) NOT NULL,
  PRIMARY KEY (disease_id, organ_symptom_id)
);

CREATE TABLE IF NOT EXISTS policy
(
  id SERIAL PRIMARY KEY,
  treatment_verb_id INTEGER REFERENCES category(id) NOT NULL,
  organ_id INTEGER REFERENCES category(id),
  tcm_organ_id INTEGER REFERENCES category(id),
  CONSTRAINT unique_policy UNIQUE(treatment_verb_id, organ_id, tcm_organ_id)
);

CREATE TABLE IF NOT EXISTS herb_effect
(
  herb_id INTEGER REFERENCES herb(id) NOT NULL,
  policy_id INTEGER REFERENCES policy(id) NOT NULL,
  PRIMARY KEY (herb_id, policy_id)
);

CREATE TABLE IF NOT EXISTS recipe_effect
(
  recipe_id INTEGER REFERENCES recipe(id) NOT NULL,
  policy_id INTEGER REFERENCES policy(id) NOT NULL,
  PRIMARY KEY (recipe_id, policy_id)
);

CREATE TABLE IF NOT EXISTS herb_meridian
(
  herb_id INTEGER REFERENCES herb(id) NOT NULL,
  meridian_id INTEGER REFERENCES category(id) NOT NULL,
  PRIMARY KEY (herb_id, meridian_id)
);

CREATE TABLE IF NOT EXISTS disease_policy
(
  disease_id INTEGER REFERENCES disease(id) NOT NULL,
  policy_id INTEGER REFERENCES policy(id) NOT NULL,
  priority INTEGER NULL,
  PRIMARY KEY (disease_id, policy_id)
);

CREATE TABLE IF NOT EXISTS pathology
(
  id SERIAL PRIMARY KEY,
  organ_id INTEGER REFERENCES category(id),
  tcm_organ_id INTEGER REFERENCES category(id),
  description_id INTEGER REFERENCES category(id) NOT NULL,
  CONSTRAINT unique_pathology UNIQUE(organ_id, tcm_organ_id, description_id)
);
--this tb can describe "心肾不交", so rival_organ tb is not needed

CREATE TABLE IF NOT EXISTS disease_pathology
(
  disease_id INTEGER REFERENCES disease(id) NOT NULL,
  pathology_id INTEGER REFERENCES pathology(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS nickname
(
  id SERIAL PRIMARY KEY,
  table_name INTEGER REFERENCES category(id) NOT NULL,
  name varchar(255) NOT NULL,
  pinyin varchar(100),
  id_in_table INTEGER NOT NULL,
  language_id INTEGER REFERENCES category(id) NOT NULL,
  CONSTRAINT unique_nickname UNIQUE(table_name, name)
);


/*
--TEMPLATE

CREATE TABLE IF NOT EXISTS foo
(
  id SERIAL PRIMARY KEY,
  name varchar(150),
  name_zh varchar(255),
  source_id INTEGER REFERENCES source(id),
  pinyin varchar(100),
  PRIMARY KEY (user_id, user2_id),
  symptom_id INTEGER REFERENCES category(id),
  CONSTRAINT unique_organ_symptom UNIQUE(organ_id, symptom_id)
);

*/