package com.wofiporia.blog.repository;

import com.wofiporia.blog.model.Blog;
import org.springframework.data.jpa.repository.JpaRepository;

public interface BlogRepository extends JpaRepository<Blog, Long> {
} 