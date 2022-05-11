// This file is generated; DO NOT EDIT.

package proxy

var validEvents = map[string]bool{
	"*":                                            true,
	"account.application.authorized":               true,
	"account.application.deauthorized":             true,
	"account.external_account.created":             true,
	"account.external_account.deleted":             true,
	"account.external_account.updated":             true,
	"account.updated":                              true,
	"application_fee.created":                      true,
	"application_fee.refund.updated":               true,
	"application_fee.refunded":                     true,
	"balance.available":                            true,
	"billing_portal.configuration.created":         true,
	"billing_portal.configuration.updated":         true,
	"billing_portal.session.created":               true,
	"capability.updated":                           true,
	"cash_balance.funds_available":                 true,
	"charge.captured":                              true,
	"charge.dispute.closed":                        true,
	"charge.dispute.created":                       true,
	"charge.dispute.funds_reinstated":              true,
	"charge.dispute.funds_withdrawn":               true,
	"charge.dispute.updated":                       true,
	"charge.expired":                               true,
	"charge.failed":                                true,
	"charge.pending":                               true,
	"charge.refund.updated":                        true,
	"charge.refunded":                              true,
	"charge.succeeded":                             true,
	"charge.updated":                               true,
	"checkout.session.async_payment_failed":        true,
	"checkout.session.async_payment_succeeded":     true,
	"checkout.session.completed":                   true,
	"checkout.session.expired":                     true,
	"coupon.created":                               true,
	"coupon.deleted":                               true,
	"coupon.updated":                               true,
	"credit_note.created":                          true,
	"credit_note.updated":                          true,
	"credit_note.voided":                           true,
	"customer.created":                             true,
	"customer.deleted":                             true,
	"customer.discount.created":                    true,
	"customer.discount.deleted":                    true,
	"customer.discount.updated":                    true,
	"customer.source.created":                      true,
	"customer.source.deleted":                      true,
	"customer.source.expiring":                     true,
	"customer.source.updated":                      true,
	"customer.subscription.created":                true,
	"customer.subscription.deleted":                true,
	"customer.subscription.pending_update_applied": true,
	"customer.subscription.pending_update_expired": true,
	"customer.subscription.trial_will_end":         true,
	"customer.subscription.updated":                true,
	"customer.tax_id.created":                      true,
	"customer.tax_id.deleted":                      true,
	"customer.tax_id.updated":                      true,
	"customer.updated":                             true,
	"file.created":                                 true,
	"identity.verification_session.canceled":       true,
	"identity.verification_session.created":        true,
	"identity.verification_session.processing":     true,
	"identity.verification_session.redacted":       true,
	"identity.verification_session.requires_input": true,
	"identity.verification_session.verified":       true,
	"invoice.created":                              true,
	"invoice.deleted":                              true,
	"invoice.finalization_failed":                  true,
	"invoice.finalized":                            true,
	"invoice.marked_uncollectible":                 true,
	"invoice.paid":                                 true,
	"invoice.payment_action_required":              true,
	"invoice.payment_failed":                       true,
	"invoice.payment_succeeded":                    true,
	"invoice.sent":                                 true,
	"invoice.upcoming":                             true,
	"invoice.updated":                              true,
	"invoice.voided":                               true,
	"invoiceitem.created":                          true,
	"invoiceitem.deleted":                          true,
	"invoiceitem.updated":                          true,
	"issuing_authorization.created":                true,
	"issuing_authorization.request":                true,
	"issuing_authorization.updated":                true,
	"issuing_card.created":                         true,
	"issuing_card.updated":                         true,
	"issuing_cardholder.created":                   true,
	"issuing_cardholder.updated":                   true,
	"issuing_dispute.closed":                       true,
	"issuing_dispute.created":                      true,
	"issuing_dispute.funds_reinstated":             true,
	"issuing_dispute.submitted":                    true,
	"issuing_dispute.updated":                      true,
	"issuing_transaction.created":                  true,
	"issuing_transaction.updated":                  true,
	"mandate.updated":                              true,
	"order.created":                                true,
	"order.payment_failed":                         true,
	"order.payment_succeeded":                      true,
	"order.updated":                                true,
	"order_return.created":                         true,
	"payment_intent.amount_capturable_updated":     true,
	"payment_intent.canceled":                      true,
	"payment_intent.created":                       true,
	"payment_intent.partially_funded":              true,
	"payment_intent.payment_failed":                true,
	"payment_intent.processing":                    true,
	"payment_intent.requires_action":               true,
	"payment_intent.succeeded":                     true,
	"payment_link.created":                         true,
	"payment_link.updated":                         true,
	"payment_method.attached":                      true,
	"payment_method.automatically_updated":         true,
	"payment_method.detached":                      true,
	"payment_method.updated":                       true,
	"payout.canceled":                              true,
	"payout.created":                               true,
	"payout.failed":                                true,
	"payout.paid":                                  true,
	"payout.updated":                               true,
	"person.created":                               true,
	"person.deleted":                               true,
	"person.updated":                               true,
	"plan.created":                                 true,
	"plan.deleted":                                 true,
	"plan.updated":                                 true,
	"price.created":                                true,
	"price.deleted":                                true,
	"price.updated":                                true,
	"product.created":                              true,
	"product.deleted":                              true,
	"product.updated":                              true,
	"promotion_code.created":                       true,
	"promotion_code.updated":                       true,
	"quote.accepted":                               true,
	"quote.canceled":                               true,
	"quote.created":                                true,
	"quote.finalized":                              true,
	"radar.early_fraud_warning.created":            true,
	"radar.early_fraud_warning.updated":            true,
	"recipient.created":                            true,
	"recipient.deleted":                            true,
	"recipient.updated":                            true,
	"reporting.report_run.failed":                  true,
	"reporting.report_run.succeeded":               true,
	"reporting.report_type.updated":                true,
	"review.closed":                                true,
	"review.opened":                                true,
	"setup_intent.canceled":                        true,
	"setup_intent.created":                         true,
	"setup_intent.requires_action":                 true,
	"setup_intent.setup_failed":                    true,
	"setup_intent.succeeded":                       true,
	"sigma.scheduled_query_run.created":            true,
	"sku.created":                                  true,
	"sku.deleted":                                  true,
	"sku.updated":                                  true,
	"source.canceled":                              true,
	"source.chargeable":                            true,
	"source.failed":                                true,
	"source.mandate_notification":                  true,
	"source.refund_attributes_required":            true,
	"source.transaction.created":                   true,
	"source.transaction.updated":                   true,
	"subscription_schedule.aborted":                true,
	"subscription_schedule.canceled":               true,
	"subscription_schedule.completed":              true,
	"subscription_schedule.created":                true,
	"subscription_schedule.expiring":               true,
	"subscription_schedule.released":               true,
	"subscription_schedule.updated":                true,
	"tax_rate.created":                             true,
	"tax_rate.updated":                             true,
	"terminal.reader.action_failed":                true,
	"terminal.reader.action_succeeded":             true,
	"test_helpers.test_clock.advancing":            true,
	"test_helpers.test_clock.created":              true,
	"test_helpers.test_clock.deleted":              true,
	"test_helpers.test_clock.internal_failure":     true,
	"test_helpers.test_clock.ready":                true,
	"topup.canceled":                               true,
	"topup.created":                                true,
	"topup.failed":                                 true,
	"topup.reversed":                               true,
	"topup.succeeded":                              true,
	"transfer.created":                             true,
	"transfer.failed":                              true,
	"transfer.paid":                                true,
	"transfer.reversed":                            true,
	"transfer.updated":                             true,
}
