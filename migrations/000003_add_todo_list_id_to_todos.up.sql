ALTER TABLE todos
ADD COLUMN todo_list_id UUID NOT NULL,
ADD CONSTRAINT fk_todo_list
    FOREIGN KEY (todo_list_id) 
    REFERENCES todo_lists(id)
    ON DELETE CASCADE;
