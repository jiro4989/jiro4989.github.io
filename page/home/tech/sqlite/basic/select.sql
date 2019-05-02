SELECt * FROM user;
SELECt * FROM hobby;

SELECT u.id, u.name, h.name, h.description
  FROM user AS u
  JOIN hobby AS h
    ON u.id = h.user_id;

