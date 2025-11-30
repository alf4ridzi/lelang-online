/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.lelang.frontend.httpclient;

/**
 *
 * @author hunter
 */
import java.io.*;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.HashMap;
import java.util.Map;
import org.json.JSONObject;
import java.util.List;

public class HttpClient {

    private String baseUrl;
    private Map<String, String> cookies;

    public HttpClient() {
        this.baseUrl = "http://localhost:8080/api";
        this.cookies = new HashMap<>();
    }

    private String executeRequest(String endpoint, String method, JSONObject requestBody) throws IOException {
        URL url = new URL(baseUrl + endpoint);
        HttpURLConnection connection = (HttpURLConnection) url.openConnection();

        try {
            connection.setRequestMethod(method);
            connection.setRequestProperty("Content-Type", "application/json");
            connection.setRequestProperty("Accept", "application/json");

            if (!cookies.isEmpty()) {
                StringBuilder cookieHeader = new StringBuilder();
                for (Map.Entry<String, String> entry : cookies.entrySet()) {
                    if (cookieHeader.length() > 0) {
                        cookieHeader.append("; ");
                    }
                    cookieHeader.append(entry.getKey()).append("=").append(entry.getValue());
                }
                
                String cookieString = cookieHeader.toString();
                
                connection.setRequestProperty("Cookie", cookieString);
            }

            if (requestBody != null && (method.equals("POST") || method.equals("PUT"))) {
                connection.setDoOutput(true);
                try (OutputStream os = connection.getOutputStream(); OutputStreamWriter osw = new OutputStreamWriter(os, "UTF-8")) {
                    osw.write(requestBody.toString());
                    osw.flush();
                }
            }

            int responseCode = connection.getResponseCode();

            InputStream inputStream;
            if (responseCode >= 200 && responseCode < 300) {
                inputStream = connection.getInputStream();
            } else {
                inputStream = connection.getErrorStream();
            }

            storeCookies(connection);

            StringBuilder response = new StringBuilder();
            try (BufferedReader br = new BufferedReader(new InputStreamReader(inputStream, "UTF-8"))) {
                String responseLine;
                while ((responseLine = br.readLine()) != null) {
                    response.append(responseLine.trim());
                }
            }

            return response.toString();

        } finally {
            connection.disconnect();
        }
    }

    private void storeCookies(HttpURLConnection connection) {
        Map<String, List<String>> headerFields = connection.getHeaderFields();
        List<String> cookieHeaders = headerFields.get("Set-Cookie");

        if (cookieHeaders != null) {
            for (String header : cookieHeaders) {
                String cookie = header.split(";", 2)[0];
                String[] parts = cookie.split("=", 2);
                if (parts.length == 2) {
                    cookies.put(parts[0], parts[1]);
                    System.out.println("Stored cookie: " + parts[0] + "=" + parts[1]);
                }
            }
        }
    }

    public JSONObject register(String username, String name, String password, String confirmpassword, String phone) {
        try {
            JSONObject requestBody = new JSONObject();
            requestBody.put("username", username);
            requestBody.put("name", name);
            requestBody.put("password", password);
            requestBody.put("confirmpassword", confirmpassword);
            requestBody.put("phone", phone);

            String response = executeRequest("/auth/register", "POST", requestBody);
            return new JSONObject(response);

        } catch (Exception e) {
            JSONObject errorResponse = new JSONObject();
            errorResponse.put("status", false);
            errorResponse.put("message", "Registration failed: " + e.getMessage());
            return errorResponse;
        }
    }

    public JSONObject login(String username, String password) {
        try {
            JSONObject requestBody = new JSONObject();
            requestBody.put("username", username);
            requestBody.put("password", password);

            String response = executeRequest("/auth/login", "POST", requestBody);
            JSONObject jsonResponse = new JSONObject(response);

            if (jsonResponse.has("status") && jsonResponse.getBoolean("status")) {
                System.out.println("sukses login");
            }

            return jsonResponse;

        } catch (Exception e) {
            JSONObject errorResponse = new JSONObject();
            errorResponse.put("status", false);
            errorResponse.put("message", "Login failed: " + e.getMessage());
            return errorResponse;
        }
    }

    public JSONObject logout() {
        try {
            String response = executeRequest("/auth/logout", "POST", null);
            JSONObject jsonResponse = new JSONObject(response);

            if (jsonResponse.has("status") && jsonResponse.getBoolean("status")) {
                cookies.clear();
                System.out.println("Logout successful! Session cleared.");
            }

            return jsonResponse;

        } catch (Exception e) {
            JSONObject errorResponse = new JSONObject();
            errorResponse.put("status", false);
            errorResponse.put("message", "Logout failed: " + e.getMessage());
            return errorResponse;
        }
    }

    public JSONObject getUserProfile() {
        try {
            String response = executeRequest("/users/profile", "GET", null);
            return new JSONObject(response);

        } catch (Exception e) {
            JSONObject errorResponse = new JSONObject();
            errorResponse.put("status", false);
            errorResponse.put("message", "Failed to get user profile: " + e.getMessage());
            return errorResponse;
        }
    }

    public boolean isAuthenticated() {
        return !cookies.isEmpty();
    }

    public Map<String, String> getCookies() {
        return new HashMap<>(cookies);
    }

    public void clearSession() {
        cookies.clear();
    }
}
