/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package websockethandler;

/**
 *
 * @author alfaridzi
 */
import com.google.gson.JsonObject;
import com.google.gson.JsonParser;
import javax.swing.SwingUtilities;
import javax.swing.table.DefaultTableModel;
import org.java_websocket.client.WebSocketClient;
import org.java_websocket.handshake.ServerHandshake;

import java.net.URI;

public class WsClient extends WebSocketClient {

    private DefaultTableModel tableModel;

    public WsClient(URI serverUri, DefaultTableModel tableModel) {
        super(serverUri);
        this.tableModel = tableModel;
    }

    @Override
    public void onOpen(ServerHandshake handshakedata) {
        System.out.println("Connected to server!");
    }

    @Override
    public void onMessage(String message) {
        System.out.println("Received: " + message);

        JsonObject json = JsonParser.parseString(message).getAsJsonObject();

        int id = json.get("id").getAsInt();
        String owner = json.get("user").getAsJsonObject().get("name").getAsString();
        String nama = json.get("item").getAsJsonObject().get("name").getAsString();
        String deskripsi = json.get("item").getAsJsonObject().get("description").getAsString();
        int bidStart = json.get("starting_bid").getAsInt();
        
        int currentBid = json.get("current_bid").getAsInt();
        int bidCount = json.get("bid_count").getAsInt();
        
        String startTime = com.lelang.frontend.dashboard.Lib.timeAgo(json.get("start_time").getAsString());
        String endTime = com.lelang.frontend.dashboard.Lib.timeAgo(json.get("end_time").getAsString());

        SwingUtilities.invokeLater(() -> {

            boolean updated = false;

            for (int i = 0; i < tableModel.getRowCount(); i++) {

                int existingID = (int) tableModel.getValueAt(i, 0);

                if (existingID == id) {
                    tableModel.setValueAt(currentBid, i, 5);
                    tableModel.setValueAt(bidCount, i, 6);

                    updated = true;
                    break;
                }
            }

            if (!updated) {
                tableModel.addRow(new Object[]{id, owner, nama, deskripsi, bidStart, currentBid, bidCount, startTime, endTime});
            }

        });
    }

    @Override
    public void onClose(int code, String reason, boolean remote) {
        System.out.println("Connection closed: " + reason);
    }

    @Override
    public void onError(Exception ex) {
        System.out.println("Error: " + ex.getMessage());
    }
}
