-- RDBMS is postgers

-- DML

DROP TABLE IF EXISTS general_tables;
DROP TABLE IF EXISTS management_base_table_custom_fields;
DROP TABLE IF EXISTS base_table_descriptions;
DROP TABLE IF EXISTS custom_field_descriptions;

CREATE TABLE base_table_descriptions (
  base_table_id SERIAL,
  description VARCHAR(255),
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  PRIMARY KEY(base_table_id)
);

CREATE TABLE custom_field_descriptions (
  custom_field_id SERIAL,
  description VARCHAR(255) NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  PRIMARY KEY(custom_field_id)
);

CREATE TABLE management_base_table_custom_fields (
  base_table_id INTEGER NOT NULL,
  custom_field_id INTEGER NOT NULL,
  PRIMARY KEY(base_table_id, custom_field_id),
  FOREIGN KEY(base_table_id) REFERENCES base_table_descriptions(base_table_id),
  FOREIGN KEY(custom_field_id) REFERENCES custom_field_descriptions(custom_field_id)
);

CREATE TABLE general_tables (
  base_table_id INTEGER NOT NULL,
  row_num SERIAL,
  custom_field_id INTEGER NOT NULL,
  field_content TEXT NOT NULL DEFAULT '',
  PRIMARY KEY(base_table_id, row_num, custom_field_id),
  FOREIGN KEY(base_table_id) REFERENCES base_table_descriptions(base_table_id),
  FOREIGN KEY(custom_field_id) REFERENCES custom_field_descriptions(custom_field_id)
);
