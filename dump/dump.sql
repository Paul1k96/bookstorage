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
INSERT INTO `authors_books` VALUES (15, 59), (5, 64), (15, 47), (13, 21), (5, 94), (6, 72), (4, 80), (2, 35), (6, 93), (13, 85), (1, 69), (5, 21), (6, 73), (5, 77), (3, 66), (6, 23), (11, 26), (3, 95), (12, 65), (5, 82), (11, 81), (13, 15), (7, 99), (10, 76), (5, 48), (12, 53), (7, 96), (15, 15), (3, 27), (12, 46), (7, 85), (14, 1), (14, 46), (2, 82), (3, 53), (4, 9), (7, 3), (1, 60), (6, 51), (15, 70), (14, 17), (5, 27), (10, 7), (9, 63), (15, 67), (7, 51), (14, 33), (12, 99), (11, 10), (3, 10), (11, 19), (9, 86), (14, 51), (8, 63), (7, 43), (10, 55), (7, 29), (15, 19), (14, 36), (6, 43), (3, 37), (2, 32), (8, 42), (7, 66), (14, 13), (8, 16), (15, 61), (7, 63), (4, 53), (13, 35), (14, 78), (1, 92), (6, 62), (12, 58), (9, 67), (13, 47), (5, 97), (13, 37), (2, 94), (3, 76), (10, 58), (14, 47), (4, 65), (15, 68), (8, 49), (11, 40), (3, 99), (15, 44), (1, 46), (15, 45), (3, 22), (6, 35), (15, 48), (9, 31), (11, 56), (4, 18), (11, 94), (14, 6), (12, 15), (13, 70), (14, 28), (9, 29), (6, 96), (14, 96), (7, 89), (4, 62), (15, 66), (14, 7), (1, 23), (5, 81), (7, 62), (14, 64), (2, 48), (4, 89), (5, 57), (1, 58), (3, 28), (5, 2), (7, 28), (7, 81), (9, 1), (3, 84), (14, 71), (5, 70), (8, 87), (13, 13), (4, 100), (9, 21), (14, 93), (15, 17), (4, 11), (3, 92), (15, 30), (14, 23), (1, 77), (11, 93), (10, 50), (5, 19), (5, 88), (12, 95), (7, 18), (5, 16), (10, 82), (7, 31), (10, 86), (1, 76), (3, 31), (1, 8), (14, 8), (6, 9), (5, 8), (1, 100), (12, 79), (11, 92), (8, 29), (11, 18), (8, 85), (10, 14), (12, 13), (13, 36), (8, 39), (10, 72), (10, 5), (3, 39), (6, 44), (12, 34), (12, 11), (7, 7), (7, 14), (3, 60), (4, 1), (2, 37), (5, 49), (3, 9), (11, 88), (7, 38), (11, 71), (11, 89), (5, 72), (11, 59), (3, 74), (7, 34), (1, 57), (13, 44), (4, 78), (9, 33), (13, 60), (2, 4), (15, 6), (2, 83), (3, 42), (2, 17), (5, 41), (4, 22), (7, 95), (7, 77), (12, 6), (3, 64), (10, 74), (13, 75), (6, 84), (7, 54), (8, 14), (4, 5), (12, 10), (7, 41), (15, 54), (10, 87), (10, 91), (4, 34), (4, 45), (3, 88), (5, 22), (15, 50), (5, 30), (2, 43), (4, 2), (6, 59), (7, 69), (14, 75), (1, 86), (1, 90), (13, 26), (13, 39), (5, 80), (4, 54), (15, 78), (4, 55), (11, 97), (12, 83), (4, 41), (11, 45), (10, 33), (15, 65), (8, 100), (7, 68), (3, 75), (7, 42), (6, 80), (15, 52), (14, 79), (9, 25), (10, 79), (9, 24), (10, 3), (8, 52), (7, 52), (8, 2), (6, 11), (1, 26), (11, 50), (9, 38), (1, 36), (8, 71), (6, 20), (3, 61), (10, 68), (4, 90), (12, 32), (14, 49), (12, 3), (2, 98), (7, 90), (4, 12);
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
