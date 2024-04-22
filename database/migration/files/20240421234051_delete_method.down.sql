DROP FUNCTION IF EXISTS delete_record(
    tbl VARCHAR(10),
    row_id UUID,
    delete_type VARCHAR(10)
);