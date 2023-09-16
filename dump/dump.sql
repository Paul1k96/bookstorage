-- MySQL dump 10.13  Distrib 8.1.0, for Linux (x86_64)
--
-- Host: localhost    Database: librarydb
-- ------------------------------------------------------
-- Server version	8.1.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `authors`
--

DROP TABLE IF EXISTS `authors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `authors` (
  `AUTHOR_ID` int NOT NULL AUTO_INCREMENT,
  `AUTHOR_NAME` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`AUTHOR_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authors`
--

LOCK TABLES `authors` WRITE;
/*!40000 ALTER TABLE `authors` DISABLE KEYS */;
INSERT INTO `authors` VALUES (1,'Brisa Gottlieb'),(2,'Adalberto Raynor'),(3,'Josue Hintz'),(4,'Theresa Dicki'),(5,'Jensen Shanahan'),(6,'Gwen Pollich'),(7,'Reymundo Langworth'),(8,'Aubree Bayer'),(9,'Friedrich Hodkiewicz'),(10,'Ted Weimann'),(11,'Glenda Kutch'),(12,'Christy Mayert'),(13,'Pamela Jast'),(14,'Jarvis Smitham'),(15,'Sarai Jacobson');
/*!40000 ALTER TABLE `authors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `authors_books`
--

DROP TABLE IF EXISTS `authors_books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `authors_books` (
  `AUTHOR_ID` int NOT NULL,
  `BOOK_ID` int NOT NULL,
  KEY `BOOK_ID` (`BOOK_ID`),
  KEY `AUTHOR_ID` (`AUTHOR_ID`),
  CONSTRAINT `authors_books_ibfk_1` FOREIGN KEY (`BOOK_ID`) REFERENCES `books` (`BOOK_ID`),
  CONSTRAINT `authors_books_ibfk_2` FOREIGN KEY (`AUTHOR_ID`) REFERENCES `authors` (`AUTHOR_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authors_books`
--

LOCK TABLES `authors_books` WRITE;
/*!40000 ALTER TABLE `authors_books` DISABLE KEYS */;
INSERT INTO `authors_books` VALUES (1,35),(14,85),(13,40),(10,43),(10,55),(4,68),(10,72),(2,34),(3,50),(2,6),(9,46),(10,75),(2,89),(4,4),(13,36),(10,47),(3,8),(14,31),(3,21),(8,7),(7,56),(3,93),(7,95),(5,15),(11,28),(8,86),(10,90),(14,91),(4,18),(1,45),(7,78),(6,92),(8,1),(9,19),(14,84),(13,99),(1,100),(9,27),(2,53),(1,77),(14,79),(13,83),(13,5),(1,11),(11,49),(2,65),(3,88),(4,3),(9,20),(7,66),(5,25),(11,26),(6,52),(3,63),(4,70),(12,87),(3,12),(12,17),(8,29),(5,33),(12,61),(3,23),(7,42),(2,48),(11,80),(12,94),(12,24),(6,58),(6,62),(5,97),(11,2),(11,10),(2,44),(7,71),(1,81),(12,16),(4,73),(3,98),(7,67),(13,96),(9,22),(13,54),(2,57),(13,76),(6,82),(11,37),(6,60),(13,30),(13,9),(14,14),(12,74),(12,13),(1,32),(12,38),(5,41),(9,69),(14,39),(2,51),(5,59),(5,64);
/*!40000 ALTER TABLE `authors_books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `books` (
  `BOOK_ID` int NOT NULL AUTO_INCREMENT,
  `BOOK_NAME` char(150) DEFAULT NULL,
  PRIMARY KEY (`BOOK_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (1,'Be am'),(2,'Firstly from'),(3,'Work tonight'),(4,'Whose yourselves'),(5,'In band'),(6,'So Torontonian'),(7,'These yet'),(8,'Downstairs including'),(9,'Outfit annually'),(10,'What us'),(11,'Therefore your'),(12,'Work it'),(13,'Dazzle whose'),(14,'Now thing'),(15,'Yesterday none'),(16,'Was bevy'),(17,'Ours Gabonese'),(18,'Neither example'),(19,'Elsewhere clarity'),(20,'Here yourselves'),(21,'Whose purely'),(22,'Then build'),(23,'Is wildly'),(24,'Ring rarely'),(25,'Itself our'),(26,'Member write'),(27,'My this'),(28,'This beneath'),(29,'Sail I'),(30,'Weekly group'),(31,'Rather on'),(32,'So Canadian'),(33,'Were this'),(34,'My though'),(35,'Of slowly'),(36,'Herself where'),(37,'Welfare day'),(38,'Nutrition have'),(39,'Differs me'),(40,'Nobody think'),(41,'Hand otherwise'),(42,'Buckles may'),(43,'Comb someone'),(44,'Peep another'),(45,'Dull as'),(46,'Firstly weekly'),(47,'Your dig'),(48,'Whenever eye'),(49,'Occur accordingly'),(50,'Climb of'),(51,'Crowd too'),(52,'Her himself'),(53,'Together case'),(54,'Annoying fact'),(55,'Bahamian climb'),(56,'One library'),(57,'Outside its'),(58,'Tenderly whose'),(59,'Awareness annually'),(60,'Class yourselves'),(61,'Stemmed pound'),(62,'But dive'),(63,'Weekly yell'),(64,'Shall jealousy'),(65,'A dive'),(66,'Girl dream'),(67,'Whichever Madagascan'),(68,'Distinguish anything'),(69,'Case last'),(70,'On adventurous'),(71,'Soon what'),(72,'Mine whoever'),(73,'To cloud'),(74,'Rather annually'),(75,'Within Balinese'),(76,'Are neatly'),(77,'Fancy weekly'),(78,'Cackle yesterday'),(79,'Daily few'),(80,'Brace each'),(81,'Rarely whichever'),(82,'Eventually is'),(83,'Spaghetti another'),(84,'Which him'),(85,'Anthology outfit'),(86,'Someone how'),(87,'Upon hour'),(88,'That then'),(89,'Late with'),(90,'Because for'),(91,'Knit there'),(92,'Warmly as'),(93,'That tense'),(94,'He life'),(95,'We were'),(96,'There harvest'),(97,'Pacific finally'),(98,'Aristotelian though'),(99,'Space exemplified'),(100,'Mushy follow');
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES (1,0);
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-09-16 16:39:47
