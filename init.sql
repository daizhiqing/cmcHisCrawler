CREATE TABLE `cmc_his_day_kline_usd` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cmcId` varchar(45) DEFAULT NULL,
  `date` varchar(45) DEFAULT NULL,
  `openPrice` varchar(45) DEFAULT NULL,
  `highPrice` varchar(45) DEFAULT NULL,
  `lowPrice` varchar(45) DEFAULT NULL,
  `closePrice` varchar(45) DEFAULT NULL,
  `volume` varchar(45) DEFAULT NULL,
  `marketCap` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `cmc_his_week_usd` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date` varchar(45) DEFAULT NULL COMMENT '日期',
  `sortId` int(11) DEFAULT NULL COMMENT 'CMC 爬取的币种ID',
  `cmcId` varchar(45) DEFAULT NULL COMMENT '全称',
  `symbol` varchar(45) DEFAULT NULL COMMENT '币种缩写',
  `marketCap` varchar(45) DEFAULT NULL COMMENT '当前市值',
  `price` varchar(45) DEFAULT NULL COMMENT '价格美元',
  `circulatingSupply` varchar(45) DEFAULT NULL COMMENT '流通量',
  `volume24` varchar(45) DEFAULT NULL COMMENT '24小时交易额',
  `h1Change` varchar(45) DEFAULT NULL COMMENT '1小时涨幅',
  `h24Change` varchar(45) DEFAULT NULL COMMENT '24小时涨幅',
  `d7Change` varchar(45) DEFAULT NULL COMMENT '7天涨幅',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=186407 DEFAULT CHARSET=utf8mb4;