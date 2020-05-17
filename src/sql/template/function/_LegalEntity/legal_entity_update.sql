-- создание юрЛица

DROP FUNCTION IF EXISTS legal_entity_update(params JSONB);
CREATE OR REPLACE FUNCTION legal_entity_update(params JSONB)
    RETURNS JSON
    LANGUAGE plpgsql
AS
$function$

DECLARE
    legal_entityRow     legal_entity%ROWTYPE;
    checkMsg    TEXT;
    result      JSONB;
    updateValue TEXT;
    queryStr    TEXT;
    
BEGIN

    -- проверика наличия id
    checkMsg = check_required_params(params, ARRAY ['id']);
    IF checkMsg IS NOT NULL
    THEN
        RETURN checkMsg;
    END IF;

    

    if (params ->> 'id')::int = -1 then
        -- проверика наличия обязательных параметров
        checkMsg = check_required_params(params, ARRAY ['title']);
        IF checkMsg IS NOT NULL
        THEN
            RETURN checkMsg;
        END IF;
        

        EXECUTE ('INSERT INTO legal_entity (title, inn, kpp, type, address_legal, options) VALUES ($1, $2, $3, $4, $5, $6)  RETURNING *;')
		INTO legal_entityRow
		USING
			(params ->> 'title')::text,
			(params ->> 'inn')::text,
			(params ->> 'kpp')::text,
			(params ->> 'type')::text,
			(params ->> 'address_legal')::text,
			coalesce(params -> 'options', '{}')::jsonb;

    else
        updateValue = '' || update_str_from_json(params, ARRAY [
			['title', 'title', 'text'],
			['inn', 'inn', 'text'],
			['kpp', 'kpp', 'text'],
			['type', 'type', 'text'],
			['address_legal', 'address_legal', 'text'],
            ['options', 'options', 'jsonb'],
            ['deleted', 'deleted', 'bool']
            ]);

        queryStr = concat('UPDATE legal_entity SET ', updateValue, ' WHERE id=', params ->> 'id', ' RETURNING *;');

        EXECUTE (queryStr)
            INTO legal_entityRow;

        -- случай когда записи с таким id не найдено
        IF row_to_json(legal_entityRow) ->> 'id' ISNULL
        THEN
            RETURN json_build_object('ok', FALSE, 'message', 'wrong id');
        END IF;

    end if;

    result = row_to_json(legal_entityRow) :: JSONB;

    RETURN json_build_object('ok', TRUE, 'result', result);

END

$function$;