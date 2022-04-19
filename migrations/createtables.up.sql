CREATE TABLE slots (
    id integer PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
    description text NOT NULL
);

CREATE TABLE banners (
    id integer PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
    description text NOT NULL
);

CREATE TABLE rotation (
    slot_id integer NOT NULL,
    banner_id integer NOT NULL,
    CONSTRAINT fk_slot
        FOREIGN KEY(slot_id)
            REFERENCES slots(id),
    CONSTRAINT fk_banner
        FOREIGN KEY(banner_id)
            REFERENCES banners(id), 
    UNIQUE (slot_id, banner_id)
);

CREATE TABLE user_groups (
    id integer PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
    description text NOT NULL
);

CREATE TABLE stat (
    slot_id integer NOT NULL,
    banner_id integer NOT NULL,
    user_group_id integer NOT NULL,
    click integer NOT NULL DEFAULT 0,
    show integer NOT NULL DEFAULT 0,
    CONSTRAINT fk_slot
        FOREIGN KEY(slot_id)
            REFERENCES slots(id),
    CONSTRAINT fk_banner
        FOREIGN KEY(banner_id)
            REFERENCES banners(id),
    CONSTRAINT fk_user_group
        FOREIGN KEY(user_group_id)
            REFERENCES user_groups(id),
    UNIQUE (slot_id, banner_id, user_group_id)
);

INSERT INTO slots (description) VALUES
('The first slot'),
('The second slot'),
('The third slot');

INSERT INTO banners (description) VALUES
('The first banner'),
('The second banner'),
('The third banner');

INSERT INTO user_groups (description) VALUES
('The first user_group'),
('The second user_group'),
('The third user_group');

INSERT INTO rotation VALUES
(1, 1),
(1, 3),
(2, 3),
(3, 1);

INSERT INTO stat (slot_id, banner_id, user_group_id) VALUES
(1, 1, 1),
(1, 3, 2),
(1, 3, 3),
(2, 3, 1),
(2, 3, 3),
(3, 1, 2),
(3, 1, 1);
