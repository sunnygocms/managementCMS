

SET FOREIGN_KEY_CHECKS=0;



-- ----------------------------
-- Table structure for `sunny_editor`
-- ----------------------------
DROP TABLE IF EXISTS `sunny_editor`;
CREATE TABLE `sunny_editor` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编辑id',
  `username` varchar(20) NOT NULL COMMENT '编辑名称',
  `password` varchar(32) NOT NULL COMMENT '密码，使用md5',
  `power` varchar(50) NOT NULL DEFAULT '' COMMENT '权限，对那些频道有创建、编辑、删除、修改权限',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态：0为已删除，1为可用',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='编辑表';

-- ----------------------------
-- Records of sunny_eeditor 默认密码123456
-- ----------------------------
INSERT INTO `sunny_editor` VALUES ('1', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '管理员', '管理员', null, '1'); 
-- ----------------------------
-- Table structure for `sunny_navigation`
-- ----------------------------
DROP TABLE IF EXISTS `sunny_navigation`;
CREATE TABLE `sunny_navigation` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `level` int(11) NOT NULL COMMENT '级别',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父ID',
  `module` char(32) DEFAULT NULL COMMENT '那个方法',
  `action` char(32) DEFAULT NULL COMMENT '那个module',
  `name` char(32) NOT NULL COMMENT '显示名称',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序，越小越靠前',
  `display` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否显示',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of aqi_navigation
-- ----------------------------

INSERT INTO `sunny_navigation` VALUES ('1', '0', '0', '', '', '系统管理', '110',  '1');
INSERT INTO `sunny_navigation` VALUES ('2', '1', '1', 'editor', 'page', '用户管理', '0','1');
INSERT INTO `sunny_navigation` VALUES ('3', '1', '1', 'Usergroup', 'page', '用户组管理', '0', '1');
INSERT INTO `sunny_navigation` VALUES ('4', '1', '1', 'power', 'page', '权限管理', '0',  '1');
INSERT INTO `sunny_navigation` VALUES ('5', '1', '1', 'navigation', 'page', '菜单管理', '0',  '1');


-- ----------------------------
-- Table structure for `sunny_power`
-- ----------------------------
DROP TABLE IF EXISTS `sunny_power`;
CREATE TABLE `sunny_power` (
  `power_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `power_name` varchar(255) DEFAULT NULL COMMENT '权限名称',
  `controller` varchar(255) DEFAULT NULL COMMENT '权限所属的控制器',
  `action` varchar(255) DEFAULT NULL COMMENT '权限下面的action',
  PRIMARY KEY (`power_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='权限表';

-- ----------------------------
-- Records of aqi_power
-- ----------------------------
INSERT INTO `sunny_power` VALUES ('1', '频道管理：修改', 'channel', 'edit');
INSERT INTO `sunny_power` VALUES ('2', '用户：添加', 'editor', 'add');
INSERT INTO `sunny_power` VALUES ('3', '用户：修改', 'editor', 'edit');
INSERT INTO `sunny_power` VALUES ('4', '用户：浏览', 'editor', 'page');
INSERT INTO `sunny_power` VALUES ('5', '用户：删除', 'editor', 'remove');
INSERT INTO `sunny_power` VALUES ('6', '权限：添加', 'power', 'add');
INSERT INTO `sunny_power` VALUES ('7', '权限：修改', 'power', 'edit');
INSERT INTO `sunny_power` VALUES ('8', '权限：浏览', 'power', 'page');
INSERT INTO `sunny_power` VALUES ('9', '权限：删除', 'power', 'del');
INSERT INTO `sunny_power` VALUES ('10', '用户组：添加', 'usergroup', 'add');
INSERT INTO `sunny_power` VALUES ('11', '用户组：修改', 'usergroup', 'edit');
INSERT INTO `sunny_power` VALUES ('12', '用户组：浏览', 'usergroup', 'page');
INSERT INTO `sunny_power` VALUES ('13', '用户组：删除', 'usergroup', 'del');
INSERT INTO `sunny_power` VALUES ('14', '菜单管理：添加', 'navigation', 'add');
INSERT INTO `sunny_power` VALUES ('15', '菜单管理：修改', 'navigation', 'edit');
INSERT INTO `sunny_power` VALUES ('16', '菜单管理：删除', 'navigation', 'remove');
INSERT INTO `sunny_power` VALUES ('17', '菜单管理：浏览', 'navigation', 'page');



-- ----------------------------
-- Table structure for `sunny_user_and_group`
-- ----------------------------
DROP TABLE IF EXISTS `sunny_user_and_group`;
CREATE TABLE `sunny_user_and_group` (
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `user_group_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户组ID',
  PRIMARY KEY (`user_id`,`user_group_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of aqi_user_and_group
-- ----------------------------
INSERT INTO `sunny_user_and_group` VALUES ('1', '1');

-- ----------------------------
-- Table structure for `sunny_user_group`
-- ----------------------------
DROP TABLE IF EXISTS `sunny_user_group`;
CREATE TABLE `sunny_user_group` (
  `user_group_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `group_name` varchar(255) DEFAULT NULL COMMENT '用户组名称',
  `edit_id` int(11) DEFAULT NULL COMMENT '添加人ID',
  `description` varchar(255) DEFAULT NULL COMMENT '用户组描述',
  `active` int(11) DEFAULT NULL COMMENT '组是否可用1 为激活状态 0 为禁用状态',
  PRIMARY KEY (`user_group_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of aqi_user_group
-- ----------------------------
INSERT INTO `sunny_user_group` VALUES ('1', '系统管理', '1', '', '0');
INSERT INTO `sunny_user_group` VALUES ('2', '全部浏览权限', '1', '所有功能的浏览权限', '1');


-- ----------------------------
-- Table structure for `sunny_usergroup_and_power`
-- ----------------------------
DROP TABLE IF EXISTS `sunny_usergroup_and_power`;
CREATE TABLE `sunny_usergroup_and_power` (
  `power_id` int(11) NOT NULL,
  `user_group_id` int(11) NOT NULL,
  PRIMARY KEY (`power_id`,`user_group_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of aqi_usergroup_and_power
-- ----------------------------
INSERT INTO `sunny_usergroup_and_power` VALUES ('1', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('2', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('3', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('4', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('5', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('6', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('7', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('8', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('9', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('10', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('11', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('11', '3');
INSERT INTO `sunny_usergroup_and_power` VALUES ('12', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('13', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('14', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('15', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('16', '1');
INSERT INTO `sunny_usergroup_and_power` VALUES ('17', '1');

