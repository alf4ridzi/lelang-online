/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.lelang.frontend.dashboard;

import com.lelang.frontend.httpclient.HttpClient;
import org.json.JSONObject;

/**
 *
 * @author alfaridzi
 */
public class Barang {
    
    private HttpClient client;

    public Barang() {
        this.client = HttpClient.getInstance();
    }
    
    public JSONObject hapusBarang(int id) {
        try {            
            String response = client.executeRequest("/items/"+id, "DELETE", null);
            return new JSONObject(response);
        } catch (Exception e) {
            JSONObject errorResponse = new JSONObject();
            errorResponse.put("status", false);
            errorResponse.put("message", e.getMessage());
            return errorResponse;
       }
    }
    
    public JSONObject editBarang(int id, String name, String description) {
        try {
            JSONObject requestBody = new JSONObject();
            requestBody.put("name", name);
            requestBody.put("description", description);
            
            String response = client.executeRequest("/items/"+id, "PUT", requestBody);
            return new JSONObject(response);
        } catch (Exception e) {
            JSONObject errorResponse = new JSONObject();
            errorResponse.put("status", false);
            errorResponse.put("message", e.getMessage());
            return errorResponse;
       }
    }
    
    public JSONObject getBarang() {
        try {
            String response = client.executeRequest("/users/items", "GET", null);
            return new JSONObject(response);
        } catch (Exception e) {
            JSONObject errorResponse = new JSONObject();
            errorResponse.put("status", false);
            errorResponse.put("message", e.getMessage());
            return errorResponse;
       }
    }
    
    public JSONObject tambahBarang(String nama, String deskripsi) {
       JSONObject data = new JSONObject();
       
       data.put("name", nama);
       data.put("description", deskripsi);
       
       try {
            String response = client.executeRequest("/items", "POST", data);
            return new JSONObject(response);
       } catch (Exception e) {
            JSONObject errorResponse = new JSONObject();
            errorResponse.put("status", false);
            errorResponse.put("message", e.getMessage());
            return errorResponse;
       }
    }
    
    
 
}
