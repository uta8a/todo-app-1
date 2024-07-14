#!/usr/local/bin/perl
use strict;
use warnings;
use Encode;
binmode STDOUT, ':utf8';
my @out = ();
my $in = decode_utf8 <STDIN>;
my @tmp = split(//, $in);

# insert space between number and non-number
# e.g. おやつは300円まで -> おやつは 300 円まで
for my $i (0 .. ($#tmp-1)) {
    if (($tmp[$i] =~ /\d/) && ($tmp[$i+1] !~ /\d/)) {
        push @out, $tmp[$i] . ' ';
    } elsif ($tmp[$i] !~ /\d/ && $tmp[$i+1] =~ /\d/) {
        push @out, $tmp[$i] . ' ';
    } else {
        push @out, $tmp[$i];
    }
}
print join('', @out);
