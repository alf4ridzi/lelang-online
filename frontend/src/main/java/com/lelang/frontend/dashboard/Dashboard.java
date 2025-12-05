/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.lelang.frontend.dashboard;

/**
 *
 * @author hunter
 */

import com.lelang.frontend.LoginForm;
import com.lelang.frontend.httpclient.HttpClient;
import javax.swing.JOptionPane;
import org.json.JSONObject;

public class Dashboard {

    private DashboardForm form;
    private HttpClient client;
    
    public Dashboard() {
        
        this.client = HttpClient.getInstance();
        
        this.form = new DashboardForm();

        loadProfile();
        form.addLogoutListener(e -> logoout());
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
    
    private void logoout() {
        int a = JOptionPane.showConfirmDialog(form, "yakin logout?", "logout", JOptionPane.YES_NO_OPTION);
        if (a != JOptionPane.YES_OPTION) {
            return;
        }
        
        JSONObject logout = client.logout();
        
        if (logout.has("status") && logout.getBoolean("status")) {
            JOptionPane.showMessageDialog(form, "berhasil logout", "berhasil", JOptionPane.INFORMATION_MESSAGE);
            LoginForm login = new LoginForm();
            login.setVisible(true);
            form.dispose();
            return;
        }
        
        JOptionPane.showMessageDialog(form, logout.getString("message"), "logout gagal", JOptionPane.ERROR_MESSAGE);
        form.dispose();
    }
}


