import React, { ReactElement, useState } from 'react';
import { Button, Form, Modal, TextArea } from '@patternfly/react-core';
import * as yup from 'yup';

import FormMessage, { FormResponseMessage } from 'Components/PatternFly/FormMessage';
import { getAxiosErrorMessage } from 'utils/responseErrorUtils';
import { useFormik } from 'formik';
import FormLabelGroup from 'Containers/Integrations/IntegrationForm/FormLabelGroup';

export type DenyDeferralFormValues = {
    comment: string;
};

export type DenyDeferralModalProps = {
    isOpen: boolean;
    numRequestsToBeAssessed: number;
    onSendRequest: (values: DenyDeferralFormValues) => Promise<FormResponseMessage>;
    onCompleteRequest: () => void;
    onCancel: () => void;
};

const validationSchema = yup.object().shape({
    comment: yup.string().trim().required('A deferral rationale is required'),
});

function DenyDeferralModal({
    isOpen,
    numRequestsToBeAssessed,
    onSendRequest,
    onCompleteRequest,
    onCancel,
}: DenyDeferralModalProps): ReactElement {
    const [message, setMessage] = useState<FormResponseMessage>(null);
    const formik = useFormik<DenyDeferralFormValues>({
        initialValues: {
            comment: '',
        },
        validationSchema,
        onSubmit: (values: DenyDeferralFormValues) => {
            const response = onSendRequest(values);
            return response;
        },
    });

    function onHandleSubmit() {
        setMessage(null);
        formik
            .submitForm()
            .then(() => {
                setMessage(null);
                formik.resetForm();
                onCompleteRequest();
            })
            .catch((response) => {
                const error = new Error(response.message);
                setMessage({
                    message: getAxiosErrorMessage(error),
                    isError: true,
                });
            });
    }

    function onChange(value, event) {
        return formik.setFieldValue(event.target.id, value);
    }

    function onCancelHandler() {
        setMessage(null);
        onCancel();
    }

    const title = `Deny deferrals (${numRequestsToBeAssessed})`;

    return (
        <Modal
            variant="small"
            title={title}
            isOpen={isOpen}
            onClose={onCancelHandler}
            actions={[
                <Button
                    key="confirm"
                    variant="danger"
                    onClick={onHandleSubmit}
                    isLoading={formik.isSubmitting}
                    isDisabled={formik.isSubmitting}
                >
                    Submit denial
                </Button>,
                <Button
                    key="cancel"
                    variant="link"
                    onClick={onCancelHandler}
                    isDisabled={formik.isSubmitting}
                >
                    Cancel
                </Button>,
            ]}
        >
            <FormMessage message={message} />
            <Form>
                <FormLabelGroup
                    label="Denial rationale"
                    isRequired
                    fieldId="comment"
                    touched={formik.touched}
                    errors={formik.errors}
                >
                    <TextArea
                        isRequired
                        type="text"
                        id="comment"
                        value={formik.values.comment}
                        onChange={(event, value) => onChange(value, event)}
                        onBlur={formik.handleBlur}
                    />
                </FormLabelGroup>
            </Form>
        </Modal>
    );
}

export default DenyDeferralModal;
