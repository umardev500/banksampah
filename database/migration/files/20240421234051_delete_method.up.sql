CREATE OR REPLACE FUNCTION delete_record(
    tbl VARCHAR(10),
    row_id UUID,
    delete_type VARCHAR(10)
) RETURNS INT AS $$
DECLARE
    sql_query TEXT;
    rows_affected INT;
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = tbl) THEN
        RAISE EXCEPTION 'Table % does not exist.', tbl;
    END IF;

    sql_query := '';
    
    IF delete_type = 'soft' THEN
        sql_query := 'UPDATE ' || tbl || ' SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1';
    ELSIF delete_type = 'hard' THEN
        sql_query := 'DELETE FROM ' || tbl || ' WHERE id = $1';
    ELSE
        RAISE EXCEPTION 'Invalid delete type. Please specify either ''soft'' or ''hard''.';
    END IF;

    EXECUTE sql_query USING row_id;
    GET DIAGNOSTICS rows_affected = ROW_COUNT;

    RETURN rows_affected;
END
$$ LANGUAGE plpgsql;