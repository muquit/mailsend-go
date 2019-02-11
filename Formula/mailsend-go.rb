class MailsendGo < Formula
  desc "Command line tool to send mail via SMTP protocol"
  homepage "https://github.com/muquit/mailsend-go"
  url "https://github.com/muquit/mailsend-go/releases/download/v1.0.1/mailsend-go_1.0.1_mac_64-bit.tar.gz"
  version "1.0.1"
  sha256 "ac6e9ddbf2f1b6e8c1c4075611edb168598f2931bf09fbd136b56d9cc69f7bd5"

  def install
    bin.install "mailsend-go"
  end
end
