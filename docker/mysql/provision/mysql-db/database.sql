
CREATE DATABASE IF NOT EXISTS proba;

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `active` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `email`, `password`, `active`, `created_at`, `updated_at`, `deleted_at`) VALUES
(21, 'pro1234@hotmail.com', 'ero2', 0, '2021-09-26 09:55:01', '2021-09-26 09:55:01', NULL),
(26, 'pro12314@hotmail.com', 'ero2222222222222222222222222222232323232', 0, '2021-09-26 09:57:26', '2021-09-26 09:57:26', NULL),
(27, 'pro11@hotmail.com', '$2a$10$Ordd5auWnn6dbZtjtdZzaeVUNUa8ca/pnfjI6AWmtrolHvlEZyIiu', 0, '2021-10-03 11:32:25', '2021-10-03 11:32:25', NULL),
(32, 'pro12@hotmail.com', '$2a$10$A7fAWGHmO6IuL.t2zgVRA.XSJB/mSH8BnK7G4jOGki7f5qs7gLz3G', 0, '2021-10-03 12:07:59', '2021-10-03 12:07:59', NULL),
(34, 'pro13@hotmail.com', '$2a$10$Z0YZlvfhgquLGa8VhEjG7uFRNBZBAMfkrmA0TDHaJWYxL4cvIF0by', 0, '2021-10-03 12:10:34', '2021-10-03 12:10:34', NULL),
(44, 'pro18thotmail.com', '$2a$10$D6F.Edt90LACb2nIF5tRO.e6DJRMpqvCYG.6SAUj5gbqnW3CunZe6', 0, '2021-10-03 12:43:48', '2021-10-03 12:43:48', NULL),
(45, 'pro189@hotmail.com', '$2a$10$H55MkXyDPcRuy.WyvAFCgusZ2SIvfuh7u9T1VQY8URkAS6ZIbQgQ6', 0, '2021-10-03 12:45:39', '2021-10-03 12:45:39', NULL),
(46, 'pro20@hotmail.com', '$2a$10$RbJymshqw7DsDeViNblKTuj20aqRPfgyRW6Xtl29vSDXgBAsUSsjC', 0, '2021-10-03 13:32:03', '2021-10-03 13:32:03', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email_unique` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=47;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;