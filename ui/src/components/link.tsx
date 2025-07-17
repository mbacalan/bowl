import { Anchor, type AnchorProps } from "@mantine/core";
import { createLink, type LinkComponent } from "@tanstack/react-router";
import { forwardRef } from "react";

interface MantineAnchorProps extends Omit<AnchorProps, "href"> {
	// Add any additional props you want to pass to the anchor
}

const MantineLinkComponent = forwardRef<HTMLAnchorElement, MantineAnchorProps>(
	(props, ref) => {
		return <Anchor ref={ref} {...props} />;
	},
);

const CreatedLinkComponent = createLink(MantineLinkComponent);

export const Link: LinkComponent<typeof MantineLinkComponent> = (props) => {
	return <CreatedLinkComponent preload="intent" {...props} />;
};
