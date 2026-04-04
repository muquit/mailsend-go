#!/usr/bin/env ruby
########################################################################
# Generate homebrew formula for my golang based projects. It depends on
# VERSION file.  Archive patterns in bin/ directory
# The formula ruby file must be copied to:
#   https://github.com/muquit/homebrew-formulae/Formula/ directory
# Usage: gen_homebrew_formula.rb --desc "Description" [--output path]
#
# First cut - Jan-17-2026 
########################################################################

require 'optimist'
require 'fileutils'

class GenHomebrewFormula
  def initialize
    $stdout.sync = true
    $stderr.sync = true
    @topts = nil
    @project_name = nil
    @version = nil
    @homepage = nil
    @description = nil
    @output_path = nil
    @checksums = {}
  end

  def log(msg)
    t = Time.new()
    $stderr.puts "#{t}: #{msg}"
  end

  def error(msg)
    $stderr.puts "ERROR: #{msg}"
    exit 1
  end

  def parse_args
    @topts = Optimist::options do
      banner "Generate Homebrew formula for Go projects"
      banner "\nUsage: gen_homebrew_formula.rb --desc \"Description\" [--output path]"
      
      opt :desc, "Description for the formula", type: :string, required: true
      opt :output, "Output directory path (default: current directory)", type: :string, default: "."
    end

    @description = @topts[:desc]
    @output_path = File.expand_path(@topts[:output])
  end

  def detect_project_name
    # Try from directory name first
    @project_name = File.basename(Dir.pwd)
    
    # Could also parse go.mod if needed
    if File.exist?('go.mod')
      content = File.read('go.mod')
      if content =~ /module\s+github\.com\/\w+\/([^\s]+)/
        @project_name = $1
      end
    end
    
    log "Detected project name: #{@project_name}"
  end

  def read_version
    error "VERSION file not found" unless File.exist?('VERSION')
    
    @version = File.read('VERSION').strip
    error "VERSION file is empty" if @version.empty?
    
    # Remove 'v' prefix if present for version variable
    @version = @version.sub(/^v/, '')
    
    log "Detected version: #{@version}"
  end

  def detect_homepage
    # Get git remote URL
    remote_url = `git remote get-url origin 2>/dev/null`.strip
    error "Could not detect git remote URL" if remote_url.empty?
    
    # Convert SSH to HTTPS
    # git@github.com:muquit/mailsend-go.git -> https://github.com/muquit/mailsend-go
    if remote_url =~ /git@github\.com:(.+)\.git$/
      @homepage = "https://github.com/#{$1}"
    elsif remote_url =~ /https:\/\/github\.com\/(.+)\.git$/
      @homepage = "https://github.com/#{$1}"
    else
      error "Could not parse git remote URL: #{remote_url}"
    end
    
    log "Detected homepage: #{@homepage}"
  end

  def read_checksums
    checksums_file = "bin/#{@project_name}-v#{@version}-checksums.txt"
    error "Checksums file not found: #{checksums_file}" unless File.exist?(checksums_file)
    
    File.readlines(checksums_file).each do |line|
      line.strip!
      next if line.empty?
      
      # Format: <hash>  <filename> (two spaces)
      if line =~ /^([a-f0-9]{64})\s+(.+)$/
        hash = $1
        filename = $2
        @checksums[filename] = hash
      end
    end
    
    log "Read #{@checksums.size} checksums"
  end

  def validate_tarballs
    required_tarballs = [
      "#{@project_name}-v#{@version}-darwin-arm64.d.tar.gz",
      "#{@project_name}-v#{@version}-darwin-amd64.d.tar.gz"
    ]
    
    required_tarballs.each do |tarball|
      tarball_path = "bin/#{tarball}"
      error "Required tarball not found: #{tarball_path}" unless File.exist?(tarball_path)
      error "Checksum not found for: #{tarball}" unless @checksums[tarball]
    end
    
    log "Validated required tarballs"
  end

  def project_to_class_name
    # Convert project-name to ClassName
    # mailsend-go -> MailsendGo
    # clip-httpd -> ClipHttpd
    @project_name.split('-').map(&:capitalize).join
  end

  def generate_formula
    class_name = project_to_class_name
    arm64_tarball = "#{@project_name}-v#{@version}-darwin-arm64.d.tar.gz"
    amd64_tarball = "#{@project_name}-v#{@version}-darwin-amd64.d.tar.gz"
    arm64_sha = @checksums[arm64_tarball]
    amd64_sha = @checksums[amd64_tarball]
    
    formula = <<~RUBY
      class #{class_name} < Formula
        desc "#{@description}"
        homepage "#{@homepage}"
        version "#{@version}"
        
        if OS.mac? && Hardware::CPU.arm?
          url "#{@homepage}/releases/download/v#{@version}/#{arm64_tarball}"
          sha256 "#{arm64_sha}"
        elsif OS.mac? && Hardware::CPU.intel?
          url "#{@homepage}/releases/download/v#{@version}/#{amd64_tarball}"
          sha256 "#{amd64_sha}"
        end
        
        def install
          if Hardware::CPU.arm?
            bin.install "#{@project_name}-v\#{version}-darwin-arm64" => "#{@project_name}"
          else
            bin.install "#{@project_name}-v\#{version}-darwin-amd64" => "#{@project_name}"
          end
        end
      end
    RUBY
    
    # Ensure output directory exists
    FileUtils.mkdir_p(@output_path) unless Dir.exist?(@output_path)
    
    output_file = File.join(@output_path, "#{@project_name}.rb")
    File.write(output_file, formula)
    
    log "Formula written to: #{output_file}"
    puts "\nGenerated formula:"
    puts "=" * 60
    puts formula
    puts "=" * 60
  end

  def doit
    parse_args()
    detect_project_name()
    read_version()
    detect_homepage()
    read_checksums()
    validate_tarballs()
    generate_formula()
  end
end

if __FILE__ == $0
  GenHomebrewFormula.new.doit()
end
