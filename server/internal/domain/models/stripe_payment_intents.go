package models

type StripePaymentIntents struct {
	ID               string `json:"id"`
	Object           string `json:"object"`
	Amount           int    `json:"amount"`
	AmountCapturable int    `json:"amount_capturable"`
	AmountDetails    struct {
		ID  string `json:"id"`
		Tip struct {
		} `json:"tip"`
	} `json:"amount_details"`
	AmountReceived          int         `json:"amount_received"`
	Application             interface{} `json:"application"`
	ApplicationFeeAmount    interface{} `json:"application_fee_amount"`
	AutomaticPaymentMethods struct {
		AllowRedirects string `json:"allow_redirects"`
		Enabled        bool   `json:"enabled"`
	} `json:"automatic_payment_methods"`
	CanceledAt         interface{} `json:"canceled_at"`
	CancellationReason interface{} `json:"cancellation_reason"`
	CaptureMethod      string      `json:"capture_method"`
	Charges            struct {
		Object string `json:"object"`
		Data   []struct {
			ID                              string      `json:"id"`
			Object                          string      `json:"object"`
			Amount                          int         `json:"amount"`
			AmountCaptured                  int         `json:"amount_captured"`
			AmountRefunded                  int         `json:"amount_refunded"`
			Application                     interface{} `json:"application"`
			ApplicationFee                  interface{} `json:"application_fee"`
			ApplicationFeeAmount            interface{} `json:"application_fee_amount"`
			ApplicationFeesRefunded         int         `json:"application_fees_refunded"`
			ApplicationFeesRefundedCurrency string      `json:"application_fees_refunded_currency"`
			BalanceTransaction              string      `json:"balance_transaction"`
			BillingDetails                  struct {
				Address struct {
					City       interface{} `json:"city"`
					Country    interface{} `json:"country"`
					Line1      interface{} `json:"line1"`
					Line2      interface{} `json:"line2"`
					PostalCode interface{} `json:"postal_code"`
					State      interface{} `json:"state"`
				} `json:"address"`
				Email interface{} `json:"email"`
				Name  interface{} `json:"name"`
				Phone interface{} `json:"phone"`
			} `json:"billing_details"`
			CalculatedStatementDescriptor string      `json:"calculated_statement_descriptor"`
			Captured                      bool        `json:"captured"`
			CapturedAt                    int         `json:"captured_at"`
			Created                       int         `json:"created"`
			CreatedWithAPIPerformanceMode bool        `json:"created_with_api_performance_mode"`
			Currency                      string      `json:"currency"`
			Customer                      string      `json:"customer"`
			Description                   interface{} `json:"description"`
			Destination                   interface{} `json:"destination"`
			Dispute                       interface{} `json:"dispute"`
			Disputed                      bool        `json:"disputed"`
			FailureBalanceTransaction     interface{} `json:"failure_balance_transaction"`
			FailureCode                   interface{} `json:"failure_code"`
			FailureMessage                interface{} `json:"failure_message"`
			FraudDetails                  struct {
			} `json:"fraud_details"`
			Geocoding struct {
				IP         interface{} `json:"ip"`
				BillingZip struct {
					Value interface{} `json:"value"`
				} `json:"billing_zip"`
			} `json:"geocoding"`
			HasGoodFundsMtReceipt       bool        `json:"has_good_funds_mt_receipt"`
			HasSvReceipt                bool        `json:"has_sv_receipt"`
			Invoice                     interface{} `json:"invoice"`
			LinkDirectlyToPaymentIntent bool        `json:"link_directly_to_payment_intent"`
			Livemode                    bool        `json:"livemode"`
			Metadata                    struct {
			} `json:"metadata"`
			MultipleDisputes bool        `json:"multiple_disputes"`
			OnBehalfOf       interface{} `json:"on_behalf_of"`
			Order            interface{} `json:"order"`
			Outcome          struct {
				DisputeProtectionResult struct {
					IsCovered        interface{} `json:"is_covered"`
					ReasonNotCovered interface{} `json:"reason_not_covered"`
				} `json:"dispute_protection_result"`
				NetworkStatus string      `json:"network_status"`
				Reason        interface{} `json:"reason"`
				RiskLevel     string      `json:"risk_level"`
				RiskScore     int         `json:"risk_score"`
				SellerMessage string      `json:"seller_message"`
				Type          string      `json:"type"`
			} `json:"outcome"`
			OwningMerchant       string `json:"owning_merchant"`
			OwningMerchantInfo   string `json:"owning_merchant_info"`
			Paid                 bool   `json:"paid"`
			PaymentIntent        string `json:"payment_intent"`
			PaymentMethod        string `json:"payment_method"`
			PaymentMethodDetails struct {
				Card struct {
					AmountAuthorized int    `json:"amount_authorized"`
					Brand            string `json:"brand"`
					CaptureBefore    int    `json:"capture_before"`
					Checks           struct {
						AddressLine1Check      interface{} `json:"address_line1_check"`
						AddressPostalCodeCheck interface{} `json:"address_postal_code_check"`
						CvcCheck               string      `json:"cvc_check"`
					} `json:"checks"`
					Country               string `json:"country"`
					ExpMonth              int    `json:"exp_month"`
					ExpYear               int    `json:"exp_year"`
					ExtendedAuthorization struct {
						Status string `json:"status"`
					} `json:"extended_authorization"`
					Fingerprint              string `json:"fingerprint"`
					Funding                  string `json:"funding"`
					IncrementalAuthorization struct {
						Status string `json:"status"`
					} `json:"incremental_authorization"`
					Installments interface{} `json:"installments"`
					IsLink       bool        `json:"is_link"`
					Issuer       string      `json:"issuer"`
					Last4        string      `json:"last4"`
					Mandate      interface{} `json:"mandate"`
					Moto         interface{} `json:"moto"`
					Multicapture struct {
						Status string `json:"status"`
					} `json:"multicapture"`
					Network      string `json:"network"`
					NetworkToken struct {
						Used bool `json:"used"`
					} `json:"network_token"`
					Overcapture struct {
						MaximumAmountCapturable int    `json:"maximum_amount_capturable"`
						Status                  string `json:"status"`
					} `json:"overcapture"`
					ThreeDSecure interface{} `json:"three_d_secure"`
					Wallet       interface{} `json:"wallet"`
				} `json:"card"`
				Type string `json:"type"`
			} `json:"payment_method_details"`
			RadarOptions struct {
				ID string `json:"id"`
			} `json:"radar_options"`
			ReceiptEmail  interface{} `json:"receipt_email"`
			ReceiptNumber interface{} `json:"receipt_number"`
			ReceiptURL    string      `json:"receipt_url"`
			Refundable    bool        `json:"refundable"`
			Refunded      bool        `json:"refunded"`
			Review        interface{} `json:"review"`
			Session       struct {
				Browser         string      `json:"browser"`
				DataSource      string      `json:"data_source"`
				Device          string      `json:"device"`
				Platform        string      `json:"platform"`
				SessionID       interface{} `json:"session_id"`
				ShowSessionData bool        `json:"show_session_data"`
				Version         interface{} `json:"version"`
			} `json:"session"`
			Shipping                          interface{} `json:"shipping"`
			Source                            interface{} `json:"source"`
			SourceTransfer                    interface{} `json:"source_transfer"`
			StatementDescriptor               interface{} `json:"statement_descriptor"`
			StatementDescriptorSuffix         interface{} `json:"statement_descriptor_suffix"`
			Status                            string      `json:"status"`
			TransferData                      interface{} `json:"transfer_data"`
			TransferGroup                     interface{} `json:"transfer_group"`
			UseSingleBtPartialCaptureBehavior bool        `json:"use_single_bt_partial_capture_behavior"`
		} `json:"data"`
		HasMore    bool   `json:"has_more"`
		TotalCount int    `json:"total_count"`
		URL        string `json:"url"`
	} `json:"charges"`
	ClientSecret                  string      `json:"client_secret"`
	ConfirmationMethod            string      `json:"confirmation_method"`
	Created                       int         `json:"created"`
	CreatedWithAPIPerformanceMode bool        `json:"created_with_api_performance_mode"`
	Currency                      string      `json:"currency"`
	Customer                      string      `json:"customer"`
	Description                   interface{} `json:"description"`
	Invoice                       interface{} `json:"invoice"`
	LastPaymentError              interface{} `json:"last_payment_error"`
	LatestCharge                  string      `json:"latest_charge"`
	Livemode                      bool        `json:"livemode"`
	Metadata                      struct {
	} `json:"metadata"`
	NextAction                        interface{} `json:"next_action"`
	OnBehalfOf                        interface{} `json:"on_behalf_of"`
	OwningMerchant                    string      `json:"owning_merchant"`
	OwningMerchantInfo                string      `json:"owning_merchant_info"`
	PaymentMethod                     string      `json:"payment_method"`
	PaymentMethodConfigurationDetails struct {
		ID     string      `json:"id"`
		Parent interface{} `json:"parent"`
	} `json:"payment_method_configuration_details"`
	PaymentMethodOptions struct {
		Card struct {
			Installments        interface{} `json:"installments"`
			MandateOptions      interface{} `json:"mandate_options"`
			Network             interface{} `json:"network"`
			RequestThreeDSecure string      `json:"request_three_d_secure"`
		} `json:"card"`
		Link struct {
			PersistentToken interface{} `json:"persistent_token"`
		} `json:"link"`
	} `json:"payment_method_options"`
	PaymentMethodTypes        []string    `json:"payment_method_types"`
	Processing                interface{} `json:"processing"`
	ReceiptEmail              interface{} `json:"receipt_email"`
	Review                    interface{} `json:"review"`
	SetupFutureUsage          interface{} `json:"setup_future_usage"`
	Shipping                  interface{} `json:"shipping"`
	Source                    interface{} `json:"source"`
	StatementDescriptor       interface{} `json:"statement_descriptor"`
	StatementDescriptorSuffix interface{} `json:"statement_descriptor_suffix"`
	Status                    string      `json:"status"`
	TransferData              interface{} `json:"transfer_data"`
	TransferGroup             interface{} `json:"transfer_group"`
}
