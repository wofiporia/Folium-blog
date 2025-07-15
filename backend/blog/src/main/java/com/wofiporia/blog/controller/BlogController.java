package com.wofiporia.blog.controller;

import com.wofiporia.blog.model.Blog;
import com.wofiporia.blog.repository.BlogRepository;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/api/blog")
public class BlogController {
    private final BlogRepository blogRepository;

    public BlogController(BlogRepository blogRepository) {
        this.blogRepository = blogRepository;
    }

    // 获取所有博客
    @GetMapping("/all")
    public List<Blog> getAll() {
        return blogRepository.findAll();
    }

    // 获取单个博客
    @GetMapping("/{id}")
    public Blog getBlog(@PathVariable Long id) {
        Optional<Blog> blog = blogRepository.findById(id);
        return blog.orElse(null);
    }

    // 新增博客
    @PostMapping("/add")
    public Blog addBlog(@RequestBody Blog blog) {
        blog.setUploadDate(LocalDateTime.now());
        blog.setUpdateDate(LocalDateTime.now());
        return blogRepository.save(blog);
    }

    // 修改博客
    @PostMapping("/update/{id}")
    public Blog updateBlog(@PathVariable Long id, @RequestBody Blog blog) {
        Optional<Blog> optionalBlog = blogRepository.findById(id);
        if (optionalBlog.isPresent()) {
            Blog existingBlog = optionalBlog.get();
            existingBlog.setTitle(blog.getTitle());
            existingBlog.setContent(blog.getContent());
            existingBlog.setUpdateDate(LocalDateTime.now());
            return blogRepository.save(existingBlog);
        }
        return null;
    }

    // 删除博客
    @DeleteMapping("/delete/{id}")
    public void deleteBlog(@PathVariable Long id) {
        blogRepository.deleteById(id);
    }
} 