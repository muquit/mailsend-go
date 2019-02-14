class MailsendGo < Formula
  desc "Command line tool to send mail via SMTP protocol"
  homepage "https://github.com/muquit/mailsend-go"
  url "https://github.com/muquit/mailsend-go/releases/download/v1.0.2/mailsend-go_1.0.2_mac_64-bit.tar.gz"
  version "1.0.2"
  sha256 "d347ffc68a67e5b452b831ce43f7f25cf0d40b4907f8f827c4be831a2eb553b5"

  def install
    bin.install "mailsend-go"
  end
end
