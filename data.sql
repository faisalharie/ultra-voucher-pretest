SELECT a.id AS id, a.name AS name, b.name AS parent_name FROM users a
LEFT JOIN users b ON a.parent_id = b.id
ORDER BY id;