import org.apache.commons.io.IOUtils;
import org.telegram.telegrambots.api.methods.send.SendMessage;
import org.telegram.telegrambots.api.objects.Update;
import org.telegram.telegrambots.bots.TelegramLongPollingBot;
import org.telegram.telegrambots.exceptions.TelegramApiException;

import java.io.IOException;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class SmotroviBot extends TelegramLongPollingBot {

    private static final String BOT_USERNAME = "bot username";
    private static final String BOT_TOKEN = "secret token";

    @Override
    public void onUpdateReceived(Update update) {
        if (update.hasMessage() && update.getMessage().hasText()) {
            String messageText = update.getMessage().getText();

            if (messageText.startsWith("/ping")) {
                String domain = messageText.replaceFirst("/ping", "").trim();

                try {
                    try {
                        double duration = ping(domain);

                        SendMessage message = new SendMessage()
                            .setChatId(update.getMessage().getChatId())
                            .setText(String.format("%s ponged with %.2fms", domain, duration));

                            sendApiMethod(message);
                    } catch (IOException e) {
                        SendMessage message = new SendMessage()
                                .setChatId(update.getMessage().getChatId())
                                .setText(String.format("error to ping %s: io exception: %s", domain, e.getMessage()));

                        sendApiMethod(message);
                    } catch (InterruptedException ignored) {
                    }
                } catch (TelegramApiException e) {
                    e.printStackTrace();
                }
            }
        }
    }

    @Override
    public String getBotUsername() {
        return BOT_USERNAME;
    }

    @Override
    public String getBotToken() {
        return BOT_TOKEN;
    }

    private double ping(String addr) throws IOException, InterruptedException {
        Process process = Runtime.getRuntime().exec("ping -c 1 " + addr);

        int returnVal = process.waitFor();
        if (returnVal != 0) {
            throw new IOException("program ping returned with non-zero code: "
                    + IOUtils.toString(process.getErrorStream(), "UTF-8"));
        }

        String out = IOUtils.toString(process.getInputStream(), "UTF-8");

        Pattern pattern = Pattern.compile("time=(.*) ms");
        Matcher matcher = pattern.matcher(out);

        if (!matcher.find()) {
            throw new IOException("program ping returned wrong output");
        }

        return Double.parseDouble(matcher.group(1));
    }
}
