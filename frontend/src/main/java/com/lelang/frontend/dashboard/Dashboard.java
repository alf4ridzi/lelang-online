/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.lelang.frontend.dashboard;

/**
 *
 * @author hunter
 */

import com.lelang.frontend.httpclient.HttpClient;
import javax.swing.JOptionPane;
import org.json.JSONObject;

public class Dashboard {

    private HttpClient client;
    private DashboardForm form;

    public Dashboard(HttpClient client) {
        this.client = client;
        
        this.form = new DashboardForm();

        loadProfile();
        form.setVisible(true);
    }

    private void loadProfile() {
        try {
            
            JSONObject user = client.getUserProfile();
            
            JSONObject data = user.getJSONObject("data");
            JSONObject roleObj = data.getJSONObject("role");
            
            String name = data.getString("name");
            String role = roleObj.getString("role");
            
            form.setUserData(name, "", role);

        } catch (Exception e) {
            JOptionPane.showMessageDialog(null, "Gagal memuat data profil! " + e.getMessage());
        }
    }
}


