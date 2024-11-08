import React, { ReactElement } from 'react';
import {
    Flex,
    FlexItem,
    Hint,
    HintBody,
    HintFooter,
    HintTitle,
    Modal,
} from '@patternfly/react-core';

import { RequestComment } from 'types/vuln_request.proto';
import { getDateTime } from 'utils/dateUtils';

export type RequestCommentsModalProps = {
    isOpen: boolean;
    cve: string;
    comments: RequestComment[];
    onClose: () => void;
};

function RequestCommentsModal({
    isOpen,
    cve,
    comments,
    onClose,
}: RequestCommentsModalProps): ReactElement {
    return (
        <Modal variant="small" title={cve} isOpen={isOpen} onClose={onClose}>
            <Flex direction={{ default: 'columnReverse' }}>
                {comments.map((comment) => {
                    return (
                        <FlexItem spacer={{ default: 'spacerLg' }}>
                            <Hint className="pf-v5-u-p-md">
                                <HintTitle className="pf-v5-u-font-size-sm pf-v5-u-font-weight-bold">
                                    {/* @TODO: Show a more descriptive text other than just the commenter's name */}
                                    {comment.user.name}
                                </HintTitle>
                                <HintBody className="pf-v5-u-font-size-sm">
                                    {comment.message}
                                </HintBody>
                                <HintFooter className="pf-v5-u-font-size-xs">
                                    {getDateTime(comment.createdAt)}
                                </HintFooter>
                            </Hint>
                        </FlexItem>
                    );
                })}
            </Flex>
        </Modal>
    );
}

export default RequestCommentsModal;
