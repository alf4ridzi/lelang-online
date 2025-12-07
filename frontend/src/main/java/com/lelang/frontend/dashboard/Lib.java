/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.lelang.frontend.dashboard;

import java.time.Instant;
import java.time.Duration;
import java.time.ZonedDateTime;
import java.time.ZoneId;

/**
 *
 * @author alfaridzi
 */
public class Lib {

    public static String timeAgo(String rfc3339) {
        try {
            Instant instant = Instant.parse(rfc3339);
            ZonedDateTime time = instant.atZone(ZoneId.systemDefault());
            ZonedDateTime now = ZonedDateTime.now();

            Duration duration = Duration.between(time, now);

            long seconds = duration.getSeconds();

            boolean future = seconds < 0;

            seconds = Math.abs(seconds);

            if (seconds < 60) {
                return future ? seconds + " detik lagi" : seconds + " detik yang lalu";
            }

            long minutes = seconds / 60;
            if (minutes < 60) {
                return future ? minutes + " menit lagi" : minutes + " menit yang lalu";
            }

            long hours = minutes / 60;
            if (hours < 24) {
                return future ? hours + " jam lagi" : hours + " jam yang lalu";
            }

            long days = hours / 24;
            return future ? days + " hari lagi" : days + " hari yang lalu";

        } catch (Exception e) {
            return "Format RFC3339 tidak valid";
        }
    }

}
