package resolver

type RootOption func(r *Root)

func RootMutation(opts ...MutationOption) RootOption {
	return func(r *Root) {
		r.mutation = NewMutation(opts...)
	}
}

func RootQuery(opts ...QueryOption) RootOption {
	return func(r *Root) {
		r.query = NewQuery(opts...)
	}
}

func RootSubscription(opts ...SubscriptionOption) RootOption {
	return func(r *Root) {
		r.subscription = NewSubscription(opts...)
	}
}
