package com.wofiporia.blog.util;

import io.jsonwebtoken.*;
import io.jsonwebtoken.security.Keys;
import java.util.Date;
import java.security.Key;

public class JwtUtil {
    // 建议放到配置文件，这里写死
    private static final String SECRET_KEY = "wofiporia-blog-very-secret-key-1234567890";
    private static final long EXPIRATION = 24 * 60 * 60 * 1000; // 24小时
    private static final Key KEY = Keys.hmacShaKeyFor(SECRET_KEY.getBytes());

    // 生成token
    public static String generateToken(String username) {
        return Jwts.builder()
                .setSubject(username)
                .setIssuedAt(new Date())
                .setExpiration(new Date(System.currentTimeMillis() + EXPIRATION))
                .signWith(KEY, SignatureAlgorithm.HS256)
                .compact();
    }

    // 解析token，返回用户名
    public static String parseToken(String token) {
        try {
            Claims claims = Jwts.parserBuilder()
                    .setSigningKey(KEY)
                    .build()
                    .parseClaimsJws(token)
                    .getBody();
            return claims.getSubject();
        } catch (JwtException e) {
            return null;
        }
    }
} 