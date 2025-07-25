package snr.kostik.events;

import net.dv8tion.jda.api.entities.emoji.Emoji;
import net.dv8tion.jda.api.events.message.MessageReceivedEvent;
import net.dv8tion.jda.api.hooks.ListenerAdapter;

import static snr.kostik.lists.ChannelList.*;
import static snr.kostik.lists.EmoteList.*;

public class MessageEventListener extends ListenerAdapter {

    @Override
    public void onMessageReceived(MessageReceivedEvent event) {
        super.onMessageReceived(event);

        if (event.getAuthor().isBot()) return;

        if (event.getChannel().getIdLong() == botTestChannelID) {
            assert cat != null;
            event.getMessage().addReaction(cat).queue();
        }

        if (event.getChannel().getIdLong() == screenshotsChannelID) {
            event.getMessage().addReaction(Emoji.fromUnicode("\uD83D\uDC4D")).queue();
            event.getMessage().addReaction(Emoji.fromUnicode("\uD83D\uDC4E")).queue();
        }

    }
}
