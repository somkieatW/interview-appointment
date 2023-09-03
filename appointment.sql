-- MySQL dump 10.13  Distrib 8.1.0, for Linux (x86_64)
--
-- Host: localhost    Database: recruitment
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
-- Table structure for table `appointment`
--

DROP TABLE IF EXISTS `appointment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `appointment` (
  `id` varchar(45) NOT NULL,
  `topic` text NOT NULL,
  `state` varchar(20) NOT NULL,
  `description` text,
  `status` varchar(1) NOT NULL,
  `created_by` varchar(45) NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_by` varchar(45) NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  KEY `appointment_FK` (`created_by`),
  KEY `appointment_FK_1` (`updated_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `appointment`
--

LOCK TABLES `appointment` WRITE;
/*!40000 ALTER TABLE `appointment` DISABLE KEYS */;
INSERT INTO `appointment` VALUES ('a2e223eb-5141-4d74-84f1-db85c3aa5a09','นัดสัมภาษณ์งาน','Done','Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam quis velit auctor, condimentum velit eu, sagittis leo. Pellentesque nec rutrum velit, vitae interdum velit. Nulla at nunc magna. Phasellus ac nibh eget purus venenatis euismod. Nunc fermentum congue neque, nec gravida erat tempor sit amet. In pellentesque blandit diam, sit amet feugiat odio vulputate quis. Aliquam nec lectus vel lacus imperdiet egestas. Maecenas leo tellus, pharetra eget porttitor pellentesque, sagittis at neque. Phasellus elementum maximus ante sed placerat. Maecenas porttitor augue tellus, id iaculis quam sollicitudin at. Mauris ut sem auctor, congue orci in, tempor orci. Duis eget turpis varius, facilisis turpis ut, suscipit tellus. Phasellus porta, ligula id ultrices pellentesque, elit sem lobortis mi, eget imperdiet nisi magna id urna.','A','b54a5d46-5734-403c-8219-0832a01a557e','2023-08-29 10:05:43','b54a5d46-5734-403c-8219-0832a01a557e','2023-09-03 01:26:52'),('ff0863b8-604f-4dae-aceb-4833034866c8','นัดสัมภาษณ์งานวันที่ 2','In Progress','ฉันชอบกินข้าวมันไก่','A','b54a5d46-5734-403c-8219-0832a01a557e','2023-09-03 00:52:40','b54a5d46-5734-403c-8219-0832a01a557e','2023-09-03 00:52:40');
/*!40000 ALTER TABLE `appointment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `id` varchar(45) NOT NULL,
  `appointment_id` varchar(45) NOT NULL,
  `user_id` varchar(45) NOT NULL,
  `content` text NOT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  KEY `comments_FK` (`appointment_id`),
  CONSTRAINT `comments_FK` FOREIGN KEY (`appointment_id`) REFERENCES `appointment` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES ('14d4ee27-6873-47e3-8a51-ccfabf4ea440','a2e223eb-5141-4d74-84f1-db85c3aa5a09','b54a5d46-5734-403c-8219-0832a01a557e','Very Niccceeee','2023-09-02 23:27:33'),('a6c12c2c-cc49-48ff-9f3d-7e7d37506edb','a2e223eb-5141-4d74-84f1-db85c3aa5a09','b54a5d46-5734-403c-8219-0832a01a557e','Not Impress :(','2023-09-03 09:10:53'),('c3da1226-6614-4213-a2a4-007064b59325','a2e223eb-5141-4d74-84f1-db85c3aa5a09','b54a5d46-5734-403c-8219-0832a01a557e','Very Good','2023-09-02 23:24:56');
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` varchar(45) NOT NULL,
  `display_name` varchar(255) NOT NULL,
  `profile_image_url` varchar(255) DEFAULT NULL,
  `status` varchar(1) NOT NULL,
  `created_by` varchar(45) NOT NULL,
  `created_date` timestamp NOT NULL,
  `email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES ('83da9d2e-8b82-4787-946b-9fc505d3c1a4','แบทแมน','https://cdn.pic.in.th/file/picinth/Batman-Party.th.jpeg','A','83da9d2e-8b82-4787-946b-9fc505d3c1a4','2023-09-03 09:18:31','batman@mail.com'),('b54a5d46-5734-403c-8219-0832a01a557e','โรบินฮู้ด','https://cdn.pic.in.th/file/picinth/Logo-Robinhood-c.png','A','b54a5d46-5734-403c-8219-0832a01a557e','2023-08-30 23:15:03','robinhood@mail.com');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-09-03 11:38:18
