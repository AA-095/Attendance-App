-- Insert mock data into `users`
INSERT INTO users (name, email) VALUES
    ('John Doe', 'john.doe2@example.com');

-- Insert mock data into `attendances`
INSERT INTO attendances (user_id, attendance_type, time, date, year) VALUES
                                                                         (1, 1, '2023-01-05 09:00:00', '2023-01-05', 2023),
                                                                         (1, 3, '2023-01-05 11:00:00', '2023-01-05', 2023),
                                                                         (1, 4, '2023-01-05 11:15:00', '2023-01-05', 2023),
                                                                         (1, 3, '2023-01-05 15:00:00', '2023-01-05', 2023),
                                                                         (1, 4, '2023-01-05 15:15:00', '2023-01-05', 2023),
                                                                         (1, 2, '2023-01-05 17:00:00', '2023-01-05', 2023),

                                                                         (1, 1, '2023-02-07 09:00:00', '2023-02-07', 2023),
                                                                         (1, 3, '2023-02-07 12:00:00', '2023-02-07', 2023),
                                                                         (1, 4, '2023-02-07 12:30:00', '2023-02-07', 2023),
                                                                         (1, 2, '2023-02-07 17:00:00', '2023-02-07', 2023),

                                                                         (1, 1, '2023-03-06 09:00:00', '2023-03-06', 2023),
                                                                         (1, 3, '2023-03-06 10:30:00', '2023-03-06', 2023),
                                                                         (1, 4, '2023-03-06 10:45:00', '2023-03-06', 2023),
                                                                         (1, 2, '2023-03-06 17:00:00', '2023-03-06', 2023),

                                                                         (1, 1, '2023-04-08 09:00:00', '2023-04-08', 2023),
                                                                         (1, 3, '2023-04-08 12:00:00', '2023-04-08', 2023),
                                                                         (1, 4, '2023-04-08 12:15:00', '2023-04-08', 2023),
                                                                         (1, 3, '2023-04-08 14:00:00', '2023-04-08', 2023),
                                                                         (1, 4, '2023-04-08 14:15:00', '2023-04-08', 2023),
                                                                         (1, 2, '2023-04-08 17:00:00', '2023-04-08', 2023),

                                                                         (1, 1, '2023-05-09 09:00:00', '2023-05-09', 2023),
                                                                         (1, 3, '2023-05-09 11:00:00', '2023-05-09', 2023),
                                                                         (1, 4, '2023-05-09 11:15:00', '2023-05-09', 2023),
                                                                         (1, 2, '2023-05-09 17:00:00', '2023-05-09', 2023),

                                                                         (1, 1, '2023-06-10 09:00:00', '2023-06-10', 2023),
                                                                         (1, 3, '2023-06-10 12:00:00', '2023-06-10', 2023),
                                                                         (1, 4, '2023-06-10 12:30:00', '2023-06-10', 2023),
                                                                         (1, 2, '2023-06-10 17:00:00', '2023-06-10', 2023),

;
