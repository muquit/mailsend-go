class MailsendGo < Formula
  desc "Command line tool to send mail via SMTP protocol"
  homepage "https://github.com/muquit/mailsend-go"
  url "https://github.com/muquit/mailsend-go/releases/download/v1.0.2/mailsend-go_1.0.2_mac_64-bit.tar.gz"
  version "1.0.2"
  sha256 "57723695d156fd813c39bf5946a69cbbad12e01260f0a39fb0c054930e6da066"

  def install
    bin.install "mailsend-go"
  end
end
