# Select Statement

Unlike switch blocks, case statements in a select block aren't tested sequentially, and execution won't automatically fall through if none of the
criteria are met. Instead, alla channel reads and writes are considered simultaneously to see if any of them are ready: populated or closed channels in the case of reads, and channels that are not at capacity in the case of writes. If none of the channels are ready, the entire select statement blocks.

## more_channels_ready

If more channels are ready, each channel has an equal chance of being selected as all the others.