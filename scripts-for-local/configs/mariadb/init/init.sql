CREATE DATABASE IF NOT EXISTS report;

USE report;

CREATE TABLE IF NOT EXISTS `sample_data` (
  `name` varchar(50) NOT NULL,
  `data_001` varchar(50) NOT NULL,
  `data_002` varchar(50) NOT NULL,
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  PRIMARY KEY (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO
  sample_data (name, data_001, data_002)
VALUES
  ('Name001', 'aaaaa', 'bbbbb'),
  ('Name002', 'ccccc', 'ddddd'),
  ('Name003', 'eeeee', 'ffffff');

GRANT ALL PRIVILEGES ON report.* TO mariauser;

FLUSH PRIVILEGES;