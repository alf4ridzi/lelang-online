/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.lelang.frontend.dashboard;

import com.lelang.frontend.httpclient.HttpClient;
import org.json.JSONObject;
import java.util.Date;
import java.time.Instant;

/**
 *
 * @author alfaridzi
 */
public class Lelang {
    private HttpClient client;
    
    public Lelang() {
        this.client = HttpClient.getInstance();
    }
    
    public JSONObject addAuction(int itemID, Date startTime, Date endTime, int startingBid) {
        String start = startTime.toInstant().toString();
        String end = endTime.toInstant().toString();
        
        try {
            JSONObject requestBody = new JSONObject();
            requestBody.put("item_id", itemID);
            requestBody.put("start_time", start);
            requestBody.put("end_time", end);
            requestBody.put("starting_bid", startingBid);
            
            String response = client.executeRequest("/auctions/new", "POST", requestBody);
            return new JSONObject(response);
        } catch (Exception e) {
            JSONObject errorResponse = new JSONObject();
            errorResponse.put("status", false);
            errorResponse.put("message", e.getMessage());
            return errorResponse;
       }
    }
}
