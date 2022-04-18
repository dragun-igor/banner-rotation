CREATE TABLE slots (
    id integer PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
    description text NOT NULL
);

CREATE TABLE banners (
    id integer PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
    description text NOT NULL
);

CREATE TABLE slots_banners (
    id integer PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
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
    slot_banner_id integer NOT NULL,
    user_group_id integer NOT NULL,
    click integer NOT NULL DEFAULT 0,
    show integer NOT NULL DEFAULT 0,
    CONSTRAINT fk_slot_banner
        FOREIGN KEY(slot_banner_id)
            REFERENCES slots_banners(id),
    CONSTRAINT fk_user_group
        FOREIGN KEY(user_group_id)
            REFERENCES user_groups(id),
    UNIQUE (slot_banner_id, user_group_id)
);