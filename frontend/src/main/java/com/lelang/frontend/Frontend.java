/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.lelang.frontend;

import com.lelang.frontend.httpclient.HttpClient;

/**
 *
 * @author hunter
 */
public class Frontend {

    public static void main(String[] args) {
        HttpClient client = new HttpClient();
         
        LoginForm login = new LoginForm(client);
        login.setVisible(true);
    }
}
