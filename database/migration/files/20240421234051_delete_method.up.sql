CREATE OR REPLACE FUNCTION delete_record(
    tbl VARCHAR(10),
    row_id UUID,
    is_soft_delete BOOLEAN = FALSE
) RETURNS INT AS $$
DECLARE
    sql_query TEXT;
    rows_affected INT;
BEGIN
    sql_query := '';
    
    IF is_soft_delete THEN
        sql_query := 'UPDATE ' || tbl || ' SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1';
    ELSE
        sql_query := 'DELETE FROM ' || tbl || ' WHERE id = $1';
    END IF;

    EXECUTE sql_query USING row_id;
    GET DIAGNOSTICS rows_affected = ROW_COUNT;

    RETURN rows_affected;
END
$$ LANGUAGE plpgsql;