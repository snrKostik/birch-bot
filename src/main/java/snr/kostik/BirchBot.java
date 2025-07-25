package snr.kostik;

import net.dv8tion.jda.api.JDABuilder;
import net.dv8tion.jda.api.OnlineStatus;
import net.dv8tion.jda.api.entities.Activity;
import net.dv8tion.jda.api.requests.GatewayIntent;
import snr.kostik.events.*;
import snr.kostik.lists.EmoteList;

import javax.security.auth.login.LoginException;

public class BirchBot {
    public static void main(String[] args) throws LoginException {
        JDABuilder jdaBuilder = JDABuilder.createDefault(Secret.discord_token);

        jdaBuilder
                .enableIntents(GatewayIntent.MESSAGE_CONTENT, GatewayIntent.GUILD_MESSAGES, GatewayIntent.GUILD_EXPRESSIONS)
                .setStatus(OnlineStatus.DO_NOT_DISTURB)
                .setActivity(Activity.playing("Зону Березовых Домов"))
                .addEventListeners(new ReadyEventListener(), new MessageEventListener(), new EmoteList())
                .build();
    }
}