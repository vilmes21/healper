INSERT INTO author (id, name) values (1,'John Doe') ON CONFLICT (id)
DO NOTHING;

INSERT INTO author (name) values ('Mark Ted') ON CONFLICT (name)
DO NOTHING;


INSERT INTO source (
      id ,
  author_id,
  language_id,
  name 
) values (1, 1, 402, 'just stuff')
ON CONFLICT (id)
DO NOTHING;

INSERT INTO disease (
      id,
  name,
  name_zh,
  source_id,
  pinyin
) values (
    1, 
    'Diabetes',
    '糖尿病',
    1,
    'tnb'
) ON CONFLICT (id) DO NOTHING;

INSERT INTO policy (
      id,
  treatment_verb_id,
  organ_id,
  tcm_organ_id
) values (
    1, 
    601,
    117,
    117
)
ON CONFLICT (treatment_verb_id,
  organ_id,
  tcm_organ_id) DO NOTHING

insert into organ_symptom (
    id,
    organ_id,
    symptom_id
) values (
    2, 117 ,703
);