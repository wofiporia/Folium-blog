package com.wofiporia.blog.controller;

import com.wofiporia.blog.model.Admin;
import com.wofiporia.blog.repository.AdminRepository;
import com.wofiporia.blog.util.JwtUtil;
import org.springframework.web.bind.annotation.*;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.HashMap;
import java.util.Map;
import java.util.Optional;
import java.util.UUID;

@RestController
@RequestMapping("/api/admin")
public class AdminController {
    @Autowired
    private AdminRepository adminRepository;

    @PostMapping("/login")
    public Map<String, Object> login(@RequestBody Map<String, String> loginData) {
        String username = loginData.get("username");
        String password = loginData.get("password");
        Optional<Admin> adminOpt = adminRepository.findByUsername(username);
        Map<String, Object> result = new HashMap<>();
        if (adminOpt.isPresent() && adminOpt.get().getPassword().equals(password)) {
            // 登录成功，生成 JWT token
            String token = JwtUtil.generateToken(username);
            result.put("success", true);
            result.put("token", token);
        } else {
            result.put("success", false);
            result.put("message", "账号或密码错误");
        }
        return result;
    }
} 