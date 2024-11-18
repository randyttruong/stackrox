import withAuth from '../../helpers/basicAuth';
import { hasFeatureFlag } from '../../helpers/features';
import { visitViolations } from './Violations.helpers';

const filteredWorkflowSelectButton = 'button[aria-label="Filtered workflow select"]';
const getFilteredWorkflowOptionButton = (viewName) =>
    `ul[aria-label="Filtered workflow select options"] button:contains("${viewName}")`;

describe('Violations - Filtered Workflow Views', () => {
    withAuth();

    before(function () {
        if (!hasFeatureFlag('ROX_PLATFORM_COMPONENTS')) {
            this.skip();
        }
    });

    it('should render the correct filtered workflow view select options', () => {
        visitViolations();

        // should exist and display "Applications view" by default
        cy.get(filteredWorkflowSelectButton).should('exist');
        cy.get(filteredWorkflowSelectButton).should('have.text', 'Applications view');

        // show the dropdown options
        cy.get(filteredWorkflowSelectButton).click();

        // the correct options should display in the dropdown
        const options = ['Applications view', 'Platform view', 'Full view'];
        options.forEach((option, index) => {
            cy.get(
                `ul[aria-label="Filtered workflow select options"] li:nth(${index}) button span.pf-v5-c-menu__item-text`
            ).should('have.text', option);
        });
    });

    it('should filter the violations table when the "Applications view" is selected', () => {
        visitViolations();

        cy.intercept('GET', '/v1/alerts?query=*').as('getViolations');

        // should filter using the correct values for the "Applications view"
        cy.wait('@getViolations').then((interception) => {
            const queryString = interception.request.query.query;

            expect(queryString).to.contain('Entity Type:DEPLOYMENT');
            expect(queryString).to.contain('Platform Component:false');
        });
    });

    it('should filter the violations table when the "Platform view" is selected', () => {
        visitViolations();

        // select the "Platform view"
        cy.get(filteredWorkflowSelectButton).click();
        cy.get(getFilteredWorkflowOptionButton('Platform view')).click();

        cy.intercept('GET', '/v1/alerts?query=*').as('getViolations');

        // should filter using the correct values for the "Platform view"
        cy.wait('@getViolations').then((interception) => {
            const queryString = interception.request.query.query;

            expect(queryString).to.contain('Entity Type:DEPLOYMENT');
            expect(queryString).to.contain('Platform Component:true');
        });
    });

    it('should filter the violations table when the "Full view" is selected', () => {
        visitViolations();

        // select the "Full view"
        cy.get(filteredWorkflowSelectButton).click();
        cy.get(getFilteredWorkflowOptionButton('Full view')).click();

        cy.intercept('GET', '/v1/alerts?query=*').as('getViolations');

        // should filter using the correct values for the "Full view"
        cy.wait('@getViolations').then((interception) => {
            const queryString = interception.request.query.query;

            expect(queryString).to.not.contain('Entity Type:DEPLOYMENT');
            expect(queryString).to.not.contain('Platform Component:true');
            expect(queryString).to.not.contain('Platform Component:false');
        });
    });
});
