/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.lelang.frontend;

import com.lelang.frontend.httpclient.HttpClient;
import org.json.JSONObject;

/**
 *
 * @author hunter
 */
public class AuthHandler {
    
    private HttpClient client;
    
    public AuthHandler(HttpClient client) {
        this.client = client;
    }
    
    public JSONObject login(String username, String password) {
        JSONObject response = client.login(username, password);
        
        return response;
    }
    
    public JSONObject register(String name, String username, String password, String confirmpassword, String phone) {
        JSONObject response = client.register(username, name, password, confirmpassword, phone);
        
        return response;
    }
    
    public JSONObject logout() {
        JSONObject response = client.logout();
        return response;
    }
}
