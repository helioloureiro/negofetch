#! /usr/bin/env perl
use strict;
use warnings;

open(my $fd, $ARGV[0]) or die "Failed to open $ARGV[0] for reading: $!";
while (<$fd>) {
    my $line = "$_";
    chomp($line);
    $line =~ s/\`/\'/g;
    $line =~ s/\$\{c(\d)\}/\`\+c\[$1\]\+\`/g;
    print "$line\n";
}
close(FD);