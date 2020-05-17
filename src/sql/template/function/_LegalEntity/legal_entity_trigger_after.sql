-- функция триггер
DROP FUNCTION IF EXISTS legal_entity_trigger_after() CASCADE;
CREATE OR REPLACE FUNCTION legal_entity_trigger_after() RETURNS trigger AS
$$
DECLARE
        r record;
BEGIN
        

    RETURN NEW;
END;

$$ LANGUAGE plpgsql;

