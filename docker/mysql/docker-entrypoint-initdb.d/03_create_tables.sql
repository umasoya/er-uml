/*******************************************/
/* sample1 tables                          */
/*******************************************/
CREATE TABLE sample1.user (
    user_id   INT         NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID',
    user_name VARCHAR(20) NOT NULL COMMENT 'ユーザー名',

    PRIMARY KEY ( user_id )
) COMMENT 'ユーザー';

CREATE TABLE sample1.post (
    post_id INT NOT NULL AUTO_INCREMENT COMMENT '投稿ID',
    user_id INT NOT NULL COMMENT 'ユーザーID',
    content TEXT NOT NULL,

    PRIMARY KEY ( post_id )
) COMMENT '投稿';
ALTER TABLE sample1.post ADD FOREIGN KEY fk_post__user(user_id) REFERENCES sample1.user(user_id);

CREATE TABLE sample1.comment (
    comment_id INT NOT NULL AUTO_INCREMENT COMMENT 'コメントID',
    post_id INT NOT NULL COMMENT '投稿ID',
    user_id INT NOT NULL COMMENT 'ユーザーID',
    comment TEXT NOT NULL,

    PRIMARY KEY ( comment_id )
) COMMENT 'コメント';
ALTER TABLE sample1.comment ADD FOREIGN KEY fk_comment__user(user_id) REFERENCES sample1.user(user_id);
ALTER TABLE sample1.comment ADD FOREIGN KEY fk_comment__post(post_id) REFERENCES sample1.post(post_id);