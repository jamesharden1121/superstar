/**
CREATE TABLE `star_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name_zh` varchar(50) NOT NULL DEFAULT '' COMMENT '中文名',
  `name_en` varchar(50) NOT NULL DEFAULT '' COMMENT '英文名',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `birthday` varchar(50) unsigned NOT NULL DEFAULT '' COMMENT '出生日期',
  `height` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '身高，单位cm',
  `weight` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '体重，单位g',
  `club` varchar(50) NOT NULL DEFAULT '' COMMENT '俱乐部',
  `jersy` varchar(50) NOT NULL DEFAULT '' COMMENT '球衣号码以及主打位置',
  `country` varchar(50) NOT NULL DEFAULT '' COMMENT '国籍',
  `birthaddress` varchar(255) NOT NULL DEFAULT '' COMMENT '出生地',
  `feature` varchar(255) NOT NULL DEFAULT '' COMMENT '个人特点',
  `moreinfo` text COMMENT '更多介绍',
  `sys_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，默认值 0 正常，1 删除',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
package models

type StarInfo struct {
	Id           int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"id"`       // 主键ID
	NameZh       string `gorm:"column:name_zh;type:varchar(50);not null" json:"name_zh"`              // 中文名
	NameEn       string `gorm:"column:name_en;type:varchar(50);not null" json:"name_en"`              // 英文名
	Avatar       string `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`               // 头像
	Birthday     string `gorm:"column:birthday;type:varchar(50);not null" json:"birthday"`            // 出生日期
	Height       int    `gorm:"column:height;type:int(10) unsigned;not null" json:"height"`           // 身高，单位cm
	Weight       int    `gorm:"column:weight;type:int(10) unsigned;not null" json:"weight"`           // 体重，单位g
	Club         string `gorm:"column:club;type:varchar(50);not null" json:"club"`                    // 俱乐部
	Jersy        string `gorm:"column:jersy;type:varchar(50);not null" json:"jersy"`                  // 球衣号码以及主打位置
	Country      string `gorm:"column:country;type:varchar(50);not null" json:"country"`              // 国籍
	Birthaddress string `gorm:"column:birthaddress;type:varchar(255);not null" json:"birthaddress"`   // 出生地
	Feature      string `gorm:"column:feature;type:varchar(255);not null" json:"feature"`             // 个人特点
	Moreinfo     string `gorm:"column:moreinfo;type:text" json:"moreinfo"`                            // 更多介绍
	SysStatus    int8   `gorm:"column:sys_status;type:tinyint(4);not null" json:"sys_status"`         // 状态，默认值 0 正常，1 删除
	SysCreated   int    `gorm:"column:sys_created;type:int(10) unsigned;not null" json:"sys_created"` // 创建时间
	SysUpdated   int    `gorm:"column:sys_updated;type:int(10) unsigned;not null" json:"sys_updated"` // 最后修改时间
}
